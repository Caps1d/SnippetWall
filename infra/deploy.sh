#!/bin/bash


APP=snippetwall
SSH_USER=ubuntu
SSH_HOST=$(terraform output -raw public_ip)
WORK_DIR="/home/$SSH_USER/apps/$APP"
SSH_KEY="~/.ssh/keys/keypairEC2.pem"


echo "Deploying to:"
echo "Host: $SSH_HOST"
echo "User: $SSH_USER"
echo "Work Directory: $WORK_DIR"

ssh -i ~/.ssh/keys/keypairEC2.pem $SSH_USER@$SSH_HOST "mkdir -p $WORK_DIR"

ssh -i ~/.ssh/keys/keypairEC2.pem $SSH_USER@$SSH_HOST "echo $PAT | docker login ghcr.io -u Caps1d --password-stdin"

rsync -avzP -e "ssh -i ~/.ssh/keys/keypairEC2.pem" ./docker-compose.yml $SSH_USER@$SSH_HOST:$WORK_DIR
rsync -avzP -e "ssh -i ~/.ssh/keys/keypairEC2.pem" ../Caddyfile $SSH_USER@$SSH_HOST:$WORK_DIR
rsync -avzP -e "ssh -i ~/.ssh/keys/keypairEC2.pem" ../internal/models/testdata/setup.sql $SSH_USER@$SSH_HOST:$WORK_DIR


ssh -i ~/.ssh/keys/keypairEC2.pem $SSH_USER@$SSH_HOST "cd $WORK_DIR && docker compose up -d"

ssh-keygen -t rsa -b 4096 -f ./ssh_key -N ""
ssh-copy-id -i ./ssh_key.pub $SSH_USER@$SSH_HOST

SSH_PRIVATE_KEY=$(cat ./ssh_key)


gh secret set SSH_PRIVATE_KEY --body "$SSH_PRIVATE_KEY"
gh secret set SSH_HOST --body "$SSH_HOST"
gh secret set SSH_USER --body "$SSH_USER"
gh secret set WORK_DIR --body "$WORK_DIR"

rm ./ssh_key ./ssh_key.pub
