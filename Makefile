install:
	sudo mv virsh-api-server.service /etc/systemd/system/virsh-api-server.service
	sudo wget -O /Virsh-Api-Server/Virsh-Api-Server https://github.com/yoanndelattre/Virsh-Api-Server/releases/download/latest/Virsh-Api-Server-Linux-x86_64
	sudo systemctl enable virsh-api-server
	sudo systemctl start virsh-api-server
