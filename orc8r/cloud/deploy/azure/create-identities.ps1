param(
        [Parameter(Mandatory)]
        [string]$rgName,

        [Parameter(Mandatory)]
        [string]$keyVaultName,

        [Parameter(Mandatory)]
        [string]$aksName,

        [Parameter()]
        [string]$location = 'eastus'
    )

$account = az account show | ConvertFrom-Json
$subscription = $account.Id

$group = (az group list --query "[?name == '$rgName']") | ConvertFrom-Json
if (-not $group -or $group.length -eq 0) 
{
    Write-Host "Creating Resource Group"
    az group create --location $location --name $rgName 
}
else {
    Write-Host "Resource Group alread exists."
}

$keyVault = (az keyvault list --query "[?name == '$keyVaultName']") | ConvertFrom-Json
if (-not $keyVault -or  $keyVault.legnth -eq 0)
{
    Write-Host "Creating key vault..."    
    az keyvault create --location $location --name $keyVaultName --resource-group $rgName --enabled-for-deployment "true" --enabled-for-template-deployment "true"
}
else {    
    Write-Host "Key Vault already exists."
}

# Ensure the AAD application for the "server" exists and is configured. Used to setup AAD login for AKS
$serverApp = (az ad app list --query "[?displayName == '${aksName}Server']") | ConvertFrom-Json
if (-not $serverApp -or $serverApp.length -eq 0) {
	Write-Host "Creating AAD app for server auth..."
    $serverApplicationId = (az ad app create --display-name "${aksName}Server" --identifier-uris "https://${aksName}Server" --query appId -o tsv)   
	Write-Host "Setting Group Membership Claims..."
	az ad app update --id $serverApplicationId --set groupMembershipClaims=All
	Write-Host "Creating Service Principal..."
	az ad sp create --id $serverApplicationId
	# Get the service principal secret
	Write-Host "Setting credential..."
	$serverApplicationSecret = $(az ad sp credential reset --name $serverApplicationId --credential-description "AKSPassword" --query password -o tsv)
	Write-Host "Setting API permissions..."
	az ad app permission add --id $serverApplicationId --api 00000003-0000-0000-c000-000000000000 --api-permissions e1fe6dd8-ba31-4d61-89e7-88639da4683d=Scope 06da0dbc-49e2-44d2-8312-53f166ab848a=Scope 7ab1d382-f21e-4acd-a863-ba3e13f7da61=Role
	Write-Host "Granting API permissions..."
	az ad app permission grant --id $serverApplicationId --api 00000003-0000-0000-c000-000000000000
	Write-Host "Applying Admin Consent..."
    az ad app permission admin-consent --id $serverApplicationId

    az keyvault secret set --vault-name $keyVaultName --name 'ServerAppId' --value $serverApplicationId
    az keyvault secret set --vault-name $keyVaultName --name 'ServerAppSecret' --value $serverApplicationSecret
}
else
{
    Write-Host "Server App already exists."
    Write-Host $serverApp
}

# Ensure the AAD application for the "client" exists and configured. Used to setup AAD login for AKS
$clientApp = (az ad app list --query "[?displayName == '${aksName}Client']") | ConvertFrom-Json
if (-not $clientApp -or $clientApp.length -eq 0) 
{
	Write-Host "Creating AAD app for client auth..."
    $clientApplicationId = (az ad app create --display-name "${aksName}Client" --native-app --reply-urls "https://${aksName}Client" --query appId -o tsv) 

	Write-Host "Creating Service Principal..."
    az ad sp create --id $clientApplicationId
    # Get the service principal secret
	Write-Host "Setting credential..."
	$clientApplicationSecret = (az ad sp credential reset --name $clientApplicationId --credential-description "AKSPassword" --query password -o tsv)
	Write-Host "Retrieving OAuth Permissions..."
	$oAuthPermissionId = (az ad app show --id $clientApplicationId --query "oauth2Permissions[0].id" -o tsv)
	Write-Host "Setting API permissions..."
	az ad app permission add --id $clientApplicationId --api $serverApplicationId --api-permissions $oAuthPermissionId=Scope
	Write-Host "Granting API permissions..."
    az ad app permission grant --id $clientApplicationId --api $serverApplicationId
    
    az keyvault secret set --vault-name $keyVaultName --name 'ClientAppId' --value $clientApplicationId
    az keyvault secret set --vault-name $keyVaultName --name 'ClientAppSecret' --value $clientApplicationSecret
}
else {
    Write-Host "Client App already exists."
}

$existingIdentity = (az ad app list --query "[?displayName == '${aksName}Identity']") | ConvertFrom-Json
if (-not $existingIdentity -or $existingIdentity.length -eq 0)
{
    Write-Host "Creating AAD app for aks identity ..."
    $clientId = (az ad app create --display-name "${aksName}Identity" --identifier-uris "https://${aksName}Identity" --query appId -o tsv)

	Write-Host "Setting Group Membership Claims..."
	az ad app update --id $clientId --set groupMembershipClaims=All
	Write-Host "Creating Service Principal..."
    $aksSp = (az ad sp create --id $clientId) | ConvertFrom-Json

	# Get the service principal secret
    $aksSecret = $(az ad sp credential reset --name $clientId --credential-description "AKSPassword" --query password -o tsv)
	az role assignment create --role "Reader" --assignee $aksSp.objectId --scope "/subscriptions/$subscription/resourceGroups/$rgName/providers/Microsoft.KeyVault/vaults/$keyVaultName"
    az keyvault set-policy -n $keyVaultName --secret-permissions get list set --spn $clientId

    az keyvault secret set --vault-name $keyVaultName --name 'ServicePrincipalObjectId' --value $aksSp.ObjectId
    az keyvault secret set --vault-name $keyVaultName --name 'ServicePrincipalClientId' --value $clientId
    az keyvault secret set --vault-name $keyVaultName --name 'ServicePrincipalSecret' --value $aksSecret
}


