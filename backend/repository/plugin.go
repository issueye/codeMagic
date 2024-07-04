package repository

import "github.com/issueye/code_magic/backend/common/model"

type CreatePlugin struct {
	model.PluginBase
}

type ModifyPlugin struct {
	model.PluginInfo
}
