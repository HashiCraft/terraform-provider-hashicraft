clean:
	rm -rf .terraform
	rm -f .terraform.lock.hcl
	rm -f terraform.tfstate
	rm -f terraform.tfstate.backup

test_apply: clean
	terraform init
	terraform apply -auto-approve

test_plan: clean
	terraform init
	terraform plan
