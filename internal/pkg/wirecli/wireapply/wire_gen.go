// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wireapply

import (
	"io"
	"sigs.k8s.io/cli-utils/internal/pkg/apply"
	"sigs.k8s.io/cli-utils/internal/pkg/clik8s"
	"sigs.k8s.io/cli-utils/internal/pkg/resourceconfig"
	"sigs.k8s.io/cli-utils/internal/pkg/util"
	"sigs.k8s.io/cli-utils/internal/pkg/wirecli/wireconfig"
	"sigs.k8s.io/cli-utils/internal/pkg/wirecli/wiregit"
	"sigs.k8s.io/cli-utils/internal/pkg/wirecli/wirek8s"
)

// Injectors from wire.go:

func InitializeApply(resourceConfigPath clik8s.ResourceConfigPath, writer io.Writer, args util.Args) (*apply.Apply, error) {
	configFlags, err := wirek8s.NewConfigFlags(args)
	if err != nil {
		return nil, err
	}
	config, err := wirek8s.NewRestConfig(configFlags)
	if err != nil {
		return nil, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		return nil, err
	}
	restMapper, err := wirek8s.NewRestMapper(config)
	if err != nil {
		return nil, err
	}
	client, err := wirek8s.NewClient(dynamicInterface, restMapper)
	if err != nil {
		return nil, err
	}
	pluginConfig := wireconfig.NewPluginConfig()
	factory := wireconfig.NewResMapFactory(pluginConfig)
	fileSystem := wireconfig.NewFileSystem()
	transformerFactory := wireconfig.NewTransformerFactory()
	kustomizeProvider := wireconfig.NewKustomizeProvider(factory, fileSystem, transformerFactory, pluginConfig)
	rawConfigFileProvider := &resourceconfig.RawConfigFileProvider{}
	configProvider := wireconfig.NewConfigProvider(resourceConfigPath, kustomizeProvider, rawConfigFileProvider)
	resourceConfigs, err := wireconfig.NewResourceConfig(resourceConfigPath, configProvider)
	if err != nil {
		return nil, err
	}
	repository := wiregit.NewOptionalRepository(resourceConfigPath)
	commitIter := wiregit.NewOptionalCommitIter(repository)
	commit := wiregit.NewOptionalCommit(commitIter)
	applyApply := &apply.Apply{
		DynamicClient: client,
		Out:           writer,
		Resources:     resourceConfigs,
		Commit:        commit,
	}
	return applyApply, nil
}

func DoApply(resourceConfigPath clik8s.ResourceConfigPath, writer io.Writer, args util.Args) (apply.Result, error) {
	configFlags, err := wirek8s.NewConfigFlags(args)
	if err != nil {
		return apply.Result{}, err
	}
	config, err := wirek8s.NewRestConfig(configFlags)
	if err != nil {
		return apply.Result{}, err
	}
	dynamicInterface, err := wirek8s.NewDynamicClient(config)
	if err != nil {
		return apply.Result{}, err
	}
	restMapper, err := wirek8s.NewRestMapper(config)
	if err != nil {
		return apply.Result{}, err
	}
	client, err := wirek8s.NewClient(dynamicInterface, restMapper)
	if err != nil {
		return apply.Result{}, err
	}
	pluginConfig := wireconfig.NewPluginConfig()
	factory := wireconfig.NewResMapFactory(pluginConfig)
	fileSystem := wireconfig.NewFileSystem()
	transformerFactory := wireconfig.NewTransformerFactory()
	kustomizeProvider := wireconfig.NewKustomizeProvider(factory, fileSystem, transformerFactory, pluginConfig)
	rawConfigFileProvider := &resourceconfig.RawConfigFileProvider{}
	configProvider := wireconfig.NewConfigProvider(resourceConfigPath, kustomizeProvider, rawConfigFileProvider)
	resourceConfigs, err := wireconfig.NewResourceConfig(resourceConfigPath, configProvider)
	if err != nil {
		return apply.Result{}, err
	}
	repository := wiregit.NewOptionalRepository(resourceConfigPath)
	commitIter := wiregit.NewOptionalCommitIter(repository)
	commit := wiregit.NewOptionalCommit(commitIter)
	applyApply := &apply.Apply{
		DynamicClient: client,
		Out:           writer,
		Resources:     resourceConfigs,
		Commit:        commit,
	}
	result, err := NewApplyCommandResult(applyApply, writer)
	if err != nil {
		return apply.Result{}, err
	}
	return result, nil
}
