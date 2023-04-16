output "vpc_id" {
  description = "The ID of the VPC"
  value       = module.vpc.vpc_id
}

output "configure_kubectl" {
  description = "Configure kubectl: make sure you're logged in with the correct AWS profile and run the following command to update your kubeconfig"
  value       = module.eks_blueprints.configure_kubectl
}

output "eks_cluster_id" {
  description = "eks cluster id "
  value       = module.eks_blueprints.eks_cluster_id
}

output "eks_cluster_endpoint" {
  description = "eks cluster endpoint"
  value       = module.eks_blueprints.eks_cluster_endpoint
}