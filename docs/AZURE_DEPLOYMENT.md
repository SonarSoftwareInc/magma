# Orchestrator Deployment to Azure

This document lays out the steps to get Orchestrator deployed and configured in Azure. It is assumed that you have the following installed on a Windows machine (refer to DEV_ENV_SETUP.md):

- Azure CLI (az) for Windows
- Windows Subsystem for Linux (WSL)
- Azure CLI for Linux
- Kubernetes CLI (kubectl) for Windows
- PowerShell

## Azure CLI Authentication

In order to use the Azure CLI, we first need to connect and authenticate with the tenant. Run the command below replacing `TENANT_ID` with the appropriate ID of your Azure tenant and enter your credentials when prompted:

`az login --tenant TENANT_ID`

## 1. Prerequisites - Azure AD Applications and Key Vault

### Creating the Resource

There are a few items that need to be present before deployment of the pipeline.  This includes definition of both a Server application and Client application within Azure AD (AAD).  Kubernetes will acquire tokens for these applications by using secrets that will be stored in Azure Key Vault and deployed to AKS secrets during the ARM template deployment.
To create the Key Vault instance, create these AAD applications and store their secrets as Key Vault secrets, run the PowerShell script: *`orc8r\cloud\deploy\azure\create-identities.ps1`* as a user that has both Azure AD global admin right as well as at least a contributor role in the resource group where the Key Vault instance will be created.
Here is an example usage of the script with the required parameters:

`.\create-identities.ps1 -rgName myorg-prod-magma -keyVaultName myorg-prod-magma-kv-1 -aksName magma-aks`

---

*Note: If for some reason, you need to re-run the script, please make sure and delete any AAD app registrations that were created previously so that the script will attempt to re-create them.*

---

## 2. Packaging and Storing Certificates In Key Vault

From a bash shell, set the current directory to a local location of the certificates by replacing the SECRET_VALUE below with the password for the certificate and running the command.

`az keyvault secret set --name SSLKeyPassword --value SECRET_VALUE --vault-name myorg-prod-magma-01`

### Packaging the Files

After creating the application certificates using the *`create_application_certs.sh`* script as described in the [Orchestrator installation documentation](https://magma.github.io/magma/docs/orc8r/deploy_install#ssl-certificates), we need to package the resulting files in order to import them. From a bash shell, run the following commands:

`openssl pkcs12 -inkey certifier.key -in certifier.pem -export -out certifier.pfx`

`openssl pkcs12 -inkey fluentd.key -in fluentd.pem -export -out fluentd.pfx`

### Importing the Certificates

Now that the certificate files are packaged, we can import them into the Key Vault by running these commands:

`az keyvault certificate import --file .\myorg.pfx --name myorg-com --vault-name myorg-prod-magma-01`

`az keyvault certificate import --file .\certifier.pfx --name Certifier --vault-name myorg-prod-magma-01`

`az keyvault certificate import --file .\fluentd.pfx --name Fluentd --vault-name myorg-prod-magma-01`

`az keyvault secret set --name BootstrapperKey --vault-name myorg-prod-magma-01 --file bootstrapper.key`

## 3.  Deploying the Infrastructure

Run the pipeline that deploys the infrastructure to your Azure tenant via the ARM templates.

Alternatively, the ARM templates can be deployed with the Azure CLI.

To deploy the net template, you can use the command below to validate and then change `validate` to `deploy` when ready to deploy:

`az deployment group validate --template-file .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.json --parameters .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.parameters.json -g myorg-prod-magma`

The other Azure resources are deployed from the "generic" ARM template.  The generic template parameters references the Key Vault secrets that were created previously.  You must edit the parameters file and replace the subscription ids, resource group names, and Key Vault instance names to point to the instance of Key Vault that were created by the script earlier.

To deploy the generic template, use the command below to validate and then change `validate` to `deploy` when ready to deploy:

`az deployment group validate --template-file .\magma\orc8r\cloud\deploy\azure\arm_templates\gen_template.json --parameters .\magma\orc8r\cloud\deploy\azure\arm_templates\gen_template.parameters.json -g myorg-prod-magma`

Each of these should execute and build all appropriate items needed for each resource group.

## 4. Container Registry

---

*Note: This will be added to the ARM Template in the future.*

---

Run the command below to create the Azure Container Registry:

`az acr create --resource-group myorg-prod-magma --name MyOrgRegistry --sku Basic`

Then run the following command to connect the ACR to AKS:

`az aks update -n magma-aks -g myorg-prod-magma --attach-acr MyOrgRegistry`

## 5.  Kubernetes CLI Authentication

Run the command below to generate AKS administrator credentials and connect the kubectl CLI to the cluster:

`az aks get-credentials --name magma-aks --resource-group myorg-prod-magma --admin`

---

*Note: Admin credentials only recommended for non-production environments.*

---

Then run the following command to create the AKS namespace:

`kubectl create namespace magma-stage`

## 6.  Populate Helm Chart Values

A values YAML file for the Helm deployment needs to be created for your specific environment. Create a YAML file under the folder: `orc8r/cloud/helm/orc8r/`.

Add the following template to the file replacing the brackets with the appropriate values:

```yaml
imagePullSecrets: []

secrets:
  create: true
secret:
  certs: orc8r-secrets-certs
  configs:
    orc8r: orc8r-secrets-configs-orc8r
  envdir: orc8r-secrets-envdir

proxy:
  podDisruptionBudget:
    enabled: true
  image:
    repository: [container registry login server]/orc8r_proxy
    tag: "latest"
    pullPolicy: Always
  replicas: 2
  service:
    enabled: true
    legacyEnabled: true
    extraAnnotations:
      bootstrapLagacy:
        external-dns.alpha.kubernetes.io/hostname: bootstrapper-controller.[domain]
      clientcertLegacy:
        external-dns.alpha.kubernetes.io/hostname: controller.[domain],api.[domain]
    name: orc8r-bootstrap-legacy
    type: LoadBalancer
  spec:
    hostname: controller.[domain]
    http_proxy_docker_hostname: "orc8r-proxy"

controller:
  podDisruptionBudget:
    enabled: true
  image:
    repository: [container registry login server]/orc8r_controller
    tag: "latest"
    pullPolicy: Always
  replicas: 2
  spec:
    database:
      db: orc8r
      host: orc8r-postgresqldb-03.postgres.database.azure.com
      port: 5432
      user: magma_dev@orc8r-postgresqldb-03
      pass: [password]

metrics:
  imagePullSecrets: []
  metrics:
    volumes:
      prometheusData:
        volumeSpec:
          persistentVolumeClaim:
            claimName: prometheus-data-azurefile
      prometheusConfig:
        volumeSpec:
          persistentVolumeClaim:
            claimName: prometheus-cfg-azurefile

  prometheus:
    create: true
    includeOrc8rAlerts: true
    prometheusCacheHostname: orc8r-prometheus-cache.magma-stage.svc.cluster.local
    alertmanagerHostname: orc8r-alertmanager.magma-stage.svc.cluster.local

  alertmanager:
    create: true

  prometheusConfigurer:
    create: true
    image:
      repository: [container registry login server]/orc8r_prometheus-configurer
      tag: "latest"
    prometheusURL: orc8r-prometheus.magma-stage.svc.cluster.local:9090

  alertmanagerConfigurer:
    create: true
    image:
      repository: [container registry login server]/orc8r_alertmanager-configurer
      tag: "latest"
    alertmanagerURL: orc8r-alertmanager.magma-stage.svc.cluster.local:9093

  prometheusCache:
    create: true
    image:
      repository: [container registry login server]/orc8r_prometheus-cache
      tag: "latest"
    limit: 500000
  grafana:
    create: false

  userGrafana:
    image:
      repository: docker.io/grafana/grafana
      tag: 6.6.2
    create: true
    volumes:
      datasources:
        persistentVolumeClaim:
          claimName: grafana-datasources-azuredisk
      dashboardproviders:
        persistentVolumeClaim:
          claimName: grafana-providers-azuredisk
      dashboards:
        persistentVolumeClaim:
          claimName: grafana-dashboards-azuredisk
      grafanaData:
        persistentVolumeClaim:
          claimName: grafana-data-azuredisk

nms:
  enabled: true

  imagePullSecrets: []

  secret:
    certs: orc8r-secrets-certs

  magmalte:
    manifests:
      secrets: true
      deployment: true
      service: true
      rbac: false

    image:
      repository: [container registry login server]/magmalte_magmalte
      tag: "latest"
      pullPolicy: Always

    env:
      api_host: orc8r-proxy.magma-stage.svc.cluster.local:9443
      mysql_host: orc8r-mysqldbserver-01.mysql.database.azure.com
      mysql_user: magma_dev@orc8r-mysqldbserver-01
      grafana_address: orc8r-user-grafana.magma-stage.svc.cluster.local:3000
      mysql_pass: [password]

  nginx:
    manifests:
      configmap: true
      secrets: true
      deployment: true
      service: true
      rbac: false

    service:
      type: LoadBalancer
      annotations:
        external-dns.alpha.kubernetes.io/hostname: "*.nms.[domain]"

    deployment:
      spec:
        ssl_cert_name: controller.crt
        ssl_cert_key_name: controller.key

logging:
  enabled: false
```

## 7.  Pipeline Configuration

### Creating the AKS Service Connection

We need to create a service connection in Azure DevOps to allow the pipeline to complete the tasks needed in the AKS cluster.

In the Azure DevOps portal, click **Project settings** in the lower left corner of the page and then click **Service connections** in the Project Settings blade.

Click the **New service connection** button in the upper right corner of the page and select **Kubernetes** from the connection type list. Click **Next**.

Ensure the authentication method selected is **Azure Subscription** and then select the target subscription from the dropdown. You may be prompted for credentials.

Once a subscription is selected, the Cluster dropdown will populate with available AKS clusters. Select the target cluster, then select the AKS namespace that we created in the above section from the dropdown, and check the box to use cluster admin credentials.

Give the service connection a name, check the box to grant permission to all pipelines, and then click **Save**.

### Creating the Pipeline

In Azure DevOps portal, click **Pipelines** in the left side of the page and then click the **New pipeline** button.

Click **Use the classic editor** on the next page. On the following page, select **GitHub** as the source and either add or select the appropriate connection. Select the correct repository and branch to deploy from and then click the **Continue** button.

On the template page, choose **YAML**.

On the next page, give the pipeline a name, then browse the path to the *`azure-container.yml`* file and select it:

**`orc8r/cloud/deploy/azure/container-pipeline.yml`**

On the Triggers tab, you can override and disable continuous integration if you only wish to have the pipeline run when triggered manually.

Click the **Save & queue** dropdown arrow and select **Save**.

## 7.  Deployment

At this point, we should be ready to run the pipeline. Open the pipeline in Azure DevOps and click the **Run pipeline** button.

From the parameters blade, we can define some values to be used during execution of the pipeline. Ensure that everything is correct and that the appropriate stages are selected in the ***Stages to run*** area.

---

*Time saving tip: The `Build_And_Publish` stage builds and publishes the docker images to the container registry so it may not be necessary to run this stage every time you trigger a deployment.*

---
