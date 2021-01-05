build:
	go build -o terraform-provider-hashicraft_v0.1.0

install: build 
	mkdir -p ~/.terraform.d/plugins/local/hashicraft/hashicraft/0.1.0/linux_amd64
	mv terraform-provider-hashicraft_v0.1.0 ~/.terraform.d/plugins/local/hashicraft/hashicraft/0.1.0/linux_amd64

run_bot:
	docker run \
	  -it \
		-e HOST=${MINECRAFT_HOST} \
		-e PORT=${MINECRAFT_PORT} \
		-e BOT_USERNAME=${MINECRAFT_USER} \
		-e BOT_PASSWORD=${MINECRAFT_PASSWORD} \
		-e API_KEY=${API_KEY} \
		-p 3000:3000 \
	  hashicraft/manicminer:v0.1.0
