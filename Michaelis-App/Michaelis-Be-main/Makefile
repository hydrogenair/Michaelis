GOCMD = go


run:
	etcd &
	$(GOCMD) run app/usercenter/cmd/rpc/usercenter.go -f app/usercenter/cmd/rpc/etc/usercenter.yaml &
	$(GOCMD) run app/usercenter/cmd/api/usercenter.go -f app/usercenter/cmd/api/etc/usercenter.yaml&
	$(GOCMD) run app/post/cmd/rpc/postcenter.go -f app/post/cmd/rpc/etc/postcenter.yaml &
	$(GOCMD) run app/comment/cmd/rpc/comment.go -f app/comment/cmd/rpc/etc/comment.yaml &
	$(GOCMD) run app/post/cmd/api/postcenter.go -f app/post/cmd/api/etc/postcenter.yaml &
	$(GOCMD) run app/chat/cmd/api/chat.go -f app/chat/cmd/api/etc/chat.yaml &
	$(GOCMD) run app/chat/cmd/rpc/chat.go -f app/chat/cmd/rpc/etc/chat.yaml &
	$(GOCMD) run app/images/cmd/api/imagecenter.go -f app/images/cmd/api/etc/imagecenter.yaml &
	$(GOCMD) run app/images/cmd/rpc/imagecenter.go -f app/images/cmd/rpc/etc/imagecenter.yaml &