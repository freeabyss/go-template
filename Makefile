
# 常量定义
SERVICE_NAME= go-template
MAIN_GO = "./main.go"


CONFIG_LOCAL = "./config/local"
CONFIG_DEV = "./config/dev"
CONFIG_PROD = "./config/prod"


APP_GO = "./go-template"
DEPLOY_TMP_DIR = ./deploy_tmp

DEV_DOCKER_PREFIX = reg.miz.so
PROD_DOCKER_PREFIX = reg.maizuo.com


ifeq (prod, ${env})
ENV=prod
DOCKER_PREFIX=$(PROD_DOCKER_PREFIX)
USE_CONTEXT=prod
NAME_SPACE=sdyxmall
DEPLOYMENT_FILE=deployment
SVC_FILE=svc
else ifeq (dmz-prod, ${env})
ENV=prod
DOCKER_PREFIX=$(PROD_DOCKER_PREFIX)
USE_CONTEXT=dmz-prod
NAME_SPACE=sdyxmall
DEPLOYMENT_FILE=deployment
SVC_FILE=svc
else ifeq (pre, ${env})
ENV=pre
DOCKER_PREFIX=$(PROD_DOCKER_PREFIX)
USE_CONTEXT=stage
NAME_SPACE=sdyxmall
DEPLOYMENT_FILE=deployment
SVC_FILE=svc
else ifeq (dmz-stage, ${env})
ENV=pre
DOCKER_PREFIX=$(PROD_DOCKER_PREFIX)
USE_CONTEXT=dmz-stage
NAME_SPACE=sdyxmall
DEPLOYMENT_FILE=deployment
SVC_FILE=svc
else
ENV=dev
DOCKER_PREFIX=$(DEV_DOCKER_PREFIX)
USE_CONTEXT=dev
NAME_SPACE=sdyxmall
DEPLOYMENT_FILE=deployment
SVC_FILE=svc
endif


# 定义命令包

define build-linux
GOOS=linux GOARCH=amd64 go build -o ${SERVICE_NAME}
endef

define run
go run $(MAIN_GO) -conf $(CONFIG_LOCAL)
endef

define docker
GOOS=linux GOARCH=amd64 go build -o $(SERVICE_NAME)
endef

define go
go run main.go -conf config/local
endef

# 清理
.PHONY : clean
clean:
	rm -rf deploy_tmp
	rm $(SERVICE_NAME)
# ----------------------------------------
# 项目初始化
.PHONY : init
init:
	@echo creating deploy directory
	test -d $(DEPLOY_TMP_DIR) || mkdir -p $(DEPLOY_TMP_DIR)


# ----------------------------------------
# 本地运行项目
.PHONY :  go
go:
	@echo starting go project
	$(go)

# ----------------------------------------
# 构建执行文件
.PHONY : package
package:
	@echo building go file
	$(build-linux)
	@echo finished building


# ----------------------------------------
# 打包docker
.PHONY : docker
docker: init check package
	@echo building docker image  tag: ${tag} env: $(ENV)
	sed  's/#env/$(ENV)/g' config/docker/Dockerfile > Dockerfile
	docker build -t $(DOCKER_PREFIX)/jike/$(SERVICE_NAME):${tag} .
	docker push $(DOCKER_PREFIX)/jike/$(SERVICE_NAME):${tag}
	@echo image is pushed !

# ----------------------------------------
# 检查参数
.PHONY : check
check:
ifeq (""a, ${tag}a)
	@echo  miss tag value
	exit 1
else ifeq (  , ${tag})
	@echo  miss tag value
	exit 1
else ifeq (""a, ${env}a)
	@echo  miss env value
	exit 1
else ifeq (  , ${env})
	@echo  miss env value
	exit 1
else
	@echo tag and env is OK
endif
	@echo tag : ${tag} env : ${env}


# ----------------------------------------
# k8s发布
.PHONY : k8s
k8s:
	kubectl config use-context $(USE_CONTEXT) --namespace=$(NAME_SPACE)
	sed  's/#tag/${tag}/g' config/k8s/$(ENV)/$(DEPLOYMENT_FILE).yaml > $(DEPLOY_TMP_DIR)/$(DEPLOYMENT_FILE)_${tag}.yaml
	kubectl apply -f config/k8s/$(ENV)/$(SVC_FILE).yaml
	kubectl apply -f $(DEPLOY_TMP_DIR)/$(DEPLOYMENT_FILE)_${tag}.yaml
	@echo apply success!
# ----------------------------------------
# 完整发布
.PHONY : deploy
deploy: init check package docker k8s
	@echo deploy complete


