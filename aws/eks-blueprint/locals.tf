locals {

  name            = basename(path.cwd)
  region          = data.aws_region.current.name
  cluster_version = "1.26"

  vpc_cidr = "10.10.0.0/16"
  azs      = slice(data.aws_availability_zones.available.names, 0, 3)

  node_group_name = "managed-ondemand"

  tags = {
    Blueprint  = local.name
    GithubRepo = "github.com/aws-ia/terraform-aws-eks-blueprints"
  }
}
