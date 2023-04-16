# Demo Project
## Project Description
This project aims to display 3 clocks with different timezones in HTML format using a Golang application. It also provides metrics at /metrics, health check at /health in JSON format and a swagger on /swagger/index.html. Additionally, it displays build time information on the frontend.

The project uses Earthly as a CI tool for building container images for AMD64 and ARM64 and pushing them to GitHub package (ghcr.io). It also enables security scans using Trivy, Snyk, and GitHub.

To deploy to Kubernetes, this project uses Helm to deploy the Demo-Project that includes Deployment, Services, Ingress, HPA, and ServiceMonitor. ArgoCD is used to apply the Helm Demo-Project and set up Kong Ingress using AWS NLB. Other tools used in the deployment process include Karpenter provisioners with spot and consolidation enabled, Certmanager setup, GoldiLocks for auto VPA to reduce costs, and Descheduler to rebalance nodes.

For infrastructure management, this project uses Terraform on AWS to create a VPC and related needs, create an EKS cluster, and install EKS and other add-ons including amazon_eks_coredns, amazon_eks_kube_proxy, amazon_eks_vpc_cni, karpenter, aws_load_balancer_controller, and metrics_server.

## Getting Started
Use Terraform to manage the infrastructure on AWS and create a VPC, EKS cluster, and install EKS and other add-ons.

To run this application, you will need to install GitHub Actions, Earthly for the CI process, Helm and ArgoCD for deploying to Kubernetes, and Terraform for infrastructure management on AWS.

Next, deploy the Helm Demo-Project using ArgoCD and set up Kong Ingress using AWS NLB. 
Extras: Apply standalone manifests for Karpenter provisioners with spot and consolidation enabled, Certmanager Setup, GoldiLocks for auto VPA, and Descheduler to rebalance nodes.

Conclusion
This project demonstrates how to build and deploy a Golang application to Kubernetes using various tools such as Earthly, Helm, ArgoCD, and Terraform. It also provides examples of using Kubernetes add-ons such as Karpenter, Certmanager, and GoldiLocks to manage and optimize the cluster.

# Specs:

## Golang
- The application will display 3 clocks with different timezones in HTML format.
- The application will expose metrics on /metrics.
- The application will have a health check on /health that returns a 200 status code in JSON format.
- The application will expose swagger on /swagger/index.html
- The application will display build time information on the frontend.

## Continuous Integration (CI)

- The project will use Earthly to build container images for AMD64 and ARM64 architectures and push them to GitHub Packages (ghcr.io).
- The runtime image for the container will use Google's Golang distroless image.
- Security scans will be enabled using Trivy, Snyk, and GitHub.

## Continuous Deployment (CD) to Kubernetes (K8S)

- Helm will be used for deploying the demo project, including services, ingress, HPA, and ServiceMonitor.
- ArgoCD will be used to apply the Helm demo project and set up Kong Ingress using AWS NLB.
- Standalone manifests will be applied, including Karpenter provisioners with spot and consolidation enabled and Certmanager setup.
- Other tools will be used, including GoldiLocks with auto VPA to reduce costs and Descheduler to rebalance nodes.

## Infrastructure

- Terraform will be used for managing the infrastructure on AWS.
- The project will use eks-blueprints as the base and will create a VPC and related resources, an EKS cluster, and install EKS addons and other addons, including amazon_eks_coredns, amazon_eks_kube_proxy, amazon_eks_vpc_cni, karpenter, aws_load_balancer_controller, and metrics_server.