package mockauthmiddleware

//go:generate mockgen -destination=./auth.go -package=mockauthmiddleware crm-system/pkg/authmiddleware AuthMiddleware
