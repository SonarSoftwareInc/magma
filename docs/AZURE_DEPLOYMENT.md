# Orchestrator Deployment to Azure

This document lays out the steps to get Orchestrator deployed and configured in Azure. It is assumed that you have the following installed on a Windows machine (refer to DEV_ENV_SETUP.md):

- Azure CLI (az) for Windows
- Windows Subsystem for Linux (WSL)
- Azure CLI for Linux
- Kubernetes CLI (kubectl) for Windows
- PowerShell

## Azure CLI Authentication

In order to use the Azure CLI, you first need to connect and authenticate with the tenant. Run the command below replacing `TENANT_ID` with the appropriate ID of your Azure tenant and enter your credentials when prompted:

`az login --tenant TENANT_ID`

## 1. Prerequisites - Azure AD Applications and Key Vault

### Creating the Resource

There are a few items that need to be present before deployment of the pipeline.  This includes definition of both a Server application and Client application within Azure AD (AAD).  Kubernetes will acquire tokens for these applications by using secrets that will be stored in Azure Key Vault and deployed to AKS secrets during the ARM template deployment.

To create the Key Vault instance, create these AAD applications and store their secrets as Key Vault secrets, run the PowerShell script: *`orc8r\cloud\deploy\azure\create-identities.ps1`* as a user that has both Azure AD global admin right as well as at least a contributor role in the resource group where the Key Vault instance will be created.
Below, is an example usage of the script with the required parameters (ensure that the Key Vault name is globally unique due to the URLs that will be generated):

`.\create-identities.ps1 -rgName myorg-prod-magma -keyVaultName myorg-prod-magma-kv-1 -aksName magma-aks`

---

*Note: If for some reason, you need to re-run the script, please make sure and delete any AAD app registrations that were created previously so that the script will attempt to re-create them.*

---

## 2. Storing Secrets and Certificates In Key Vault

### Creating the Secrets

From a bash shell, set the current directory to a local location of the certificates by replacing the SECRET_VALUE below with the password for the certificate and running the command below.

`az keyvault secret set --name SSLKeyPassword --value SECRET_VALUE --vault-name myorg-prod-magma-01`

Run the command below replacing DB_USER with the username to be used with the controller database.

`az keyvault secret set --name ControllerDBUser --value DB_USER --vault-name myorg-prod-magma-01`

Run the command below replacing DB_PASS with the password to be used with the controller database.

`az keyvault secret set --name ControllerDBPass --value DB_PASS --vault-name myorg-prod-magma-01`

Run the command below replacing DB_USER with the username to be used with the MagmaLTE database.

`az keyvault secret set --name MagmaLTEDBUser --value DB_USER --vault-name myorg-prod-magma-01`

Run the command below replacing DB_PASS with the password to be used with the MagmaLTE database.

`az keyvault secret set --name MagmaLTEDBPass --value DB_PASS --vault-name myorg-prod-magma-01`

### Packaging the Files

After creating the application certificates using the *`create_application_certs.sh`* script as described in the [Orchestrator installation documentation](https://magma.github.io/magma/docs/orc8r/deploy_install#ssl-certificates), you need to package the resulting files in order to import them. If you're planning to use a self-signed certificate for the domain, you will want to refer to the same documentation and run the *`self_sign_certs.sh`* script as well. The second parameter you will pass to *`self_sign_certs.sh`* is the password that will be used to package the PFX file (use the same value you set the `SSLKeyPassword` secret to). Then, from a bash shell, run the following commands replacing CERT_PASSWORD with the value you set the `SSLKeyPassword` secret to:

`openssl pkcs12 -inkey certifier.key -in certifier.pem -export -out certifier.pfx -password pass:"CERT_PASSWORD"`

`openssl pkcs12 -inkey fluentd.key -in fluentd.pem -export -out fluentd.pfx -password pass:"CERT_PASSWORD"`

### Importing the Certificates

Now that the certificate files are packaged, you can import them into the Key Vault by running the commands below replacing CERT_PASSWORD with the value you set the `SSLKeyPassword` secret to. If you are using a self-signed cert, the *`self_sign_certs.sh`* script will have created the *`domain.pfx`* file used in the first command for you when you ran it in the previous step. If you already have a cert then ensure the filename is correct in the command.

`az keyvault certificate import --file .\domain.pfx --name domain --vault-name myorg-prod-magma-01 --password CERT_PASSWORD`

`az keyvault certificate import --file .\certifier.pfx --name Certifier --vault-name myorg-prod-magma-01 --password CERT_PASSWORD`

`az keyvault certificate import --file .\fluentd.pfx --name Fluentd --vault-name myorg-prod-magma-01 --password CERT_PASSWORD`

`az keyvault secret set --name BootstrapperKey --vault-name myorg-prod-magma-01 --file bootstrapper.key`

## 3.  Deploying the Infrastructure

The ARM templates can be deployed with the Azure CLI.

To deploy the net template, you can use the command below to validate and then change `validate` to `deploy` when ready to deploy (replace `myorg-prod-magma` with the appropriate resource group):

`az deployment group validate --template-file .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.json --parameters .\magma\orc8r\cloud\deploy\azure\arm_templates\net_template.parameters.json -g myorg-prod-magma`

The other Azure resources are deployed from the "generic" ARM template.  The generic template parameters references the Key Vault secrets that were created previously.  You must edit the parameters file and replace the subscription ids, resource group names, and Key Vault instance names to point to the instance of Key Vault that were created by the script earlier.

To deploy the generic template, use the command below to validate and then change `validate` to `deploy` when ready to deploy (replace `myorg-prod-magma` with the appropriate resource group).

---

*`Imporant`: The values for `postgresql_db_server_name` and `mysql_db_server_name` in the `gen_template.parameter.json` file need to be globally unique so you will need to open the file and set these prior to deploying the template.*

---

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

A values YAML file for the Helm deployment needs to be created for your specific environment. Create a YAML file in a separate private repository for the appropriate deployment stage, naming it as such (replacing `stage` with the actual name of the stage i.e. "prod" or "stage"):

`vals-stage.yml`

Add the following template to the file replacing the curly braces and their contents with the appropriate values:

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
    repository: {container registry login server}/orc8r_proxy
    tag: "latest"
    pullPolicy: Always
  replicas: 2
  service:
    enabled: true
    legacyEnabled: true
    extraAnnotations:
      bootstrapLagacy:
        external-dns.alpha.kubernetes.io/hostname: bootstrapper-controller.{myorg.com}
      clientcertLegacy:
        external-dns.alpha.kubernetes.io/hostname: controller.{myorg.com},api.{myorg.com}
    name: orc8r-bootstrap-legacy
    type: LoadBalancer
  spec:
    hostname: controller.{myorg.com}
    http_proxy_docker_hostname: "orc8r-proxy"

controller:
  podDisruptionBudget:
    enabled: true
  image:
    repository: {container registry login server}/orc8r_controller
    tag: "latest"
    pullPolicy: Always
  replicas: 2
  spec:
    database:
      db: orc8r
      host: {value of postgresql_db_server_name from ARM template}.postgres.database.azure.com
      port: 5432
      user:
      pass:

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
    prometheusCacheHostname: orc8r-prometheus-cache.{AKS namespace}.svc.cluster.local
    alertmanagerHostname: orc8r-alertmanager.{AKS namespace}.svc.cluster.local

  alertmanager:
    create: true

  prometheusConfigurer:
    create: true
    image:
      repository: {container registry login server}/orc8r_prometheus-configurer
      tag: "latest"
    prometheusURL: orc8r-prometheus.{AKS namespace}.svc.cluster.local:9090

  alertmanagerConfigurer:
    create: true
    image:
      repository: {container registry login server}/orc8r_alertmanager-configurer
      tag: "latest"
    alertmanagerURL: orc8r-alertmanager.{AKS namespace}.svc.cluster.local:9093

  prometheusCache:
    create: true
    image:
      repository: {container registry login server}/orc8r_prometheus-cache
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
  enabled: false

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
      repository: {container registry login server}/magmalte_magmalte
      tag: "latest"
      pullPolicy: Always

    env:
      api_host: orc8r-proxy.{AKS namespace}.svc.cluster.local:9443
      mysql_host: {value of mysql_db_server_name from ARM template}.mysql.database.azure.com
      mysql_user:
      grafana_address: orc8r-user-grafana.{AKS namespace}.svc.cluster.local:3000
      mysql_pass:

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
        external-dns.alpha.kubernetes.io/hostname: "*.nms.{myorg.com}"

    deployment:
      spec:
        ssl_cert_name: controller.crt
        ssl_cert_key_name: controller.key

logging:
  enabled: false
```

## 7.  Pipeline Configuration

### Creating the AKS Service Connection

You need to create a service connection in Azure DevOps to allow the pipeline to complete the tasks needed in the AKS cluster.

In the Azure DevOps portal, click **Project settings** in the lower left corner of the page and then click **Service connections** in the Project Settings blade.

Click the **New service connection** button in the upper right corner of the page and select **Kubernetes** from the connection type list. Click **Next**.

Ensure the authentication method selected is **Azure Subscription** and then select the target subscription from the dropdown. You may be prompted for credentials.

Once a subscription is selected, the Cluster dropdown will populate with available AKS clusters. Select the target cluster, then select the AKS namespace that you created in the above section from the dropdown, and check the box to use cluster admin credentials.

Give the service connection a name, check the box to grant permission to all pipelines, and then click **Save**.

### Updating the YAML

Open the pipeline YAML file for editing which is located here:

*`orc8r/cloud/deploy/azure/container-pipeline.yml`*

In the section below, replace the curly braces and their contents to point to the private repository you created in the [Populate Helm Chart Values](#6-populate-helm-chart-values) section:

```yaml
resources:
  repositories:
  - repository: magma_config
    type: github
    endpoint: {GitHub service connection}
    name: {MyOrg/magma_config}
```

---

*Note: The `repository` value is the name used to reference the repository internally within the pipeline. To avoid updating the aks-pipeline.yml file, do not change this value.`

The `GitHub service connection` can be the same one that was already created to allow Azure DevOps to access the magma code base.*

---

### Creating the Pipeline

In Azure DevOps portal, click **Pipelines** in the left side of the page and then click the **New pipeline** button.

Click **Use the classic editor** on the next page. On the following page, select **GitHub** as the source and either add or select the appropriate connection. Select the correct repository and branch to deploy from and then click the **Continue** button.

On the template page, choose **YAML**.

On the next page, give the pipeline a name, then browse the path to the *`azure-container.yml`* file and select it:

**`orc8r/cloud/deploy/azure/container-pipeline.yml`**

On the Triggers tab, you can override and disable continuous integration if you only wish to have the pipeline run when triggered manually.

Click the **Save & queue** dropdown arrow and select **Save**.

## 7.  Deployment

At this point, you should be ready to run the pipeline. For a fresh install, you will need to run it twice - first with nms disabled in the vals file you created for the intented environment in the [Populate Helm Chart Values](#6-populate-helm-chart-values) section and after you import the `admin_operator.pfx` certificate to Key Vault, you will run it again with nms enabled. For subsequent runs, nms will be left enabled and the pipeline only needs to be ran once. The steps below assume a fresh install is being carried out.

### 1. Run the Pipeline with NMS Disabled

Open the pipeline in Azure DevOps and click the **Run pipeline** button.

From the parameters blade, you will define some values to be used during execution of the pipeline. Update the parameter values for the intended environment ensuring that everything is correct and that the appropriate stages are selected in the ***Stages to run*** area.

---

*Note: Checking the `Enable system diagnostics` should only be done for troubleshooting issues with the pipeline as it may cause some tasks in the pipeline to appear to fail when actually successful due to the excessive logging.*

---

### 2. Import the admin_operator Certificate

After creating the steps in the [Orchestrator installation documentation](https://magma.github.io/magma/docs/orc8r/deploy_install#creating-an-admin-user) for creating an admin user, you should have the `admin_operator.pfx` downloaded on your machine. You will import that file to Key Vault with this command (replacing CERT_PASSWORD with the same value you set the `SSLKeyPassword` secret to):

`az keyvault certificate import --file .\admin_operator.pfx --name admin-operator --vault-name myorg-prod-magma-01 --password CERT_PASSWORD`

### 3. Run the Pipeline with NMS Enabled

Next, open the vals file for the intended environment you created in the [Populate Helm Chart Values](#6-populate-helm-chart-values) section and modify this section, setting the value for the `enabled` property to true:

```YAML
nms:
  enabled: true
```

After you save, commit, and push the changes to the private repository, you can then re-run the pipeline and the nms containers will be created.

---

*Time saving tip: The `Build_And_Publish` stage builds and publishes the docker images to the container registry so it may not be necessary to run this stage every time you trigger a deployment if the images haven't changed.*

---
