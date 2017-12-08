// Code generated by cdpgen. DO NOT EDIT.

package runtime

// AwaitPromiseArgs represents the arguments for AwaitPromise in the Runtime domain.
type AwaitPromiseArgs struct {
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`           // Identifier of the promise.
	ReturnByValue   *bool          `json:"returnByValue,omitempty"`   // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview *bool          `json:"generatePreview,omitempty"` // Whether preview should be generated for the result.
}

// NewAwaitPromiseArgs initializes AwaitPromiseArgs with the required arguments.
func NewAwaitPromiseArgs(promiseObjectID RemoteObjectID) *AwaitPromiseArgs {
	args := new(AwaitPromiseArgs)
	args.PromiseObjectID = promiseObjectID
	return args
}

// SetReturnByValue sets the ReturnByValue optional argument. Whether the result is expected to be a JSON object that should be sent by value.
func (a *AwaitPromiseArgs) SetReturnByValue(returnByValue bool) *AwaitPromiseArgs {
	a.ReturnByValue = &returnByValue
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument. Whether preview should be generated for the result.
func (a *AwaitPromiseArgs) SetGeneratePreview(generatePreview bool) *AwaitPromiseArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// AwaitPromiseReply represents the return values for AwaitPromise in the Runtime domain.
type AwaitPromiseReply struct {
	Result           RemoteObject      `json:"result"`                     // Promise result. Will contain rejected value if promise was rejected.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details if stack strace is available.
}

// CallFunctionOnArgs represents the arguments for CallFunctionOn in the Runtime domain.
type CallFunctionOnArgs struct {
	FunctionDeclaration string          `json:"functionDeclaration"`     // Declaration of the function to call.
	ObjectID            *RemoteObjectID `json:"objectId,omitempty"`      // Identifier of the object to call function on. Either objectId or executionContextId should be specified.
	Arguments           []CallArgument  `json:"arguments,omitempty"`     // Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Silent              *bool           `json:"silent,omitempty"`        // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	ReturnByValue       *bool           `json:"returnByValue,omitempty"` // Whether the result is expected to be a JSON object which should be sent by value.
	// GeneratePreview Whether preview should be generated for the result.
	//
	// Note: This property is experimental.
	GeneratePreview    *bool               `json:"generatePreview,omitempty"`
	UserGesture        *bool               `json:"userGesture,omitempty"`        // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise       *bool               `json:"awaitPromise,omitempty"`       // Whether execution should `await` for resulting value and return once awaited promise is resolved.
	ExecutionContextID *ExecutionContextID `json:"executionContextId,omitempty"` // Specifies execution context which global object will be used to call function on. Either executionContextId or objectId should be specified.
	ObjectGroup        *string             `json:"objectGroup,omitempty"`        // Symbolic group name that can be used to release multiple objects. If objectGroup is not specified and objectId is, objectGroup will be inherited from object.
}

// NewCallFunctionOnArgs initializes CallFunctionOnArgs with the required arguments.
func NewCallFunctionOnArgs(functionDeclaration string) *CallFunctionOnArgs {
	args := new(CallFunctionOnArgs)
	args.FunctionDeclaration = functionDeclaration
	return args
}

// SetObjectID sets the ObjectID optional argument. Identifier of the object to call function on. Either objectId or executionContextId should
// be specified.
func (a *CallFunctionOnArgs) SetObjectID(objectID RemoteObjectID) *CallFunctionOnArgs {
	a.ObjectID = &objectID
	return a
}

// SetArguments sets the Arguments optional argument. Call arguments. All call arguments must belong to the same JavaScript world as the target
// object.
func (a *CallFunctionOnArgs) SetArguments(arguments []CallArgument) *CallFunctionOnArgs {
	a.Arguments = arguments
	return a
}

// SetSilent sets the Silent optional argument. In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (a *CallFunctionOnArgs) SetSilent(silent bool) *CallFunctionOnArgs {
	a.Silent = &silent
	return a
}

// SetReturnByValue sets the ReturnByValue optional argument. Whether the result is expected to be a JSON object which should be sent by value.
func (a *CallFunctionOnArgs) SetReturnByValue(returnByValue bool) *CallFunctionOnArgs {
	a.ReturnByValue = &returnByValue
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument. Whether preview should be generated for the result.
//
// Note: This property is experimental.
func (a *CallFunctionOnArgs) SetGeneratePreview(generatePreview bool) *CallFunctionOnArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// SetUserGesture sets the UserGesture optional argument. Whether execution should be treated as initiated by user in the UI.
func (a *CallFunctionOnArgs) SetUserGesture(userGesture bool) *CallFunctionOnArgs {
	a.UserGesture = &userGesture
	return a
}

// SetAwaitPromise sets the AwaitPromise optional argument. Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (a *CallFunctionOnArgs) SetAwaitPromise(awaitPromise bool) *CallFunctionOnArgs {
	a.AwaitPromise = &awaitPromise
	return a
}

// SetExecutionContextID sets the ExecutionContextID optional argument. Specifies execution context which global object will be used to call function on. Either
// executionContextId or objectId should be specified.
func (a *CallFunctionOnArgs) SetExecutionContextID(executionContextID ExecutionContextID) *CallFunctionOnArgs {
	a.ExecutionContextID = &executionContextID
	return a
}

// SetObjectGroup sets the ObjectGroup optional argument. Symbolic group name that can be used to release multiple objects. If objectGroup is not
// specified and objectId is, objectGroup will be inherited from object.
func (a *CallFunctionOnArgs) SetObjectGroup(objectGroup string) *CallFunctionOnArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// CallFunctionOnReply represents the return values for CallFunctionOn in the Runtime domain.
type CallFunctionOnReply struct {
	Result           RemoteObject      `json:"result"`                     // Call result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// CompileScriptArgs represents the arguments for CompileScript in the Runtime domain.
type CompileScriptArgs struct {
	Expression         string              `json:"expression"`                   // Expression to compile.
	SourceURL          string              `json:"sourceURL"`                    // Source url to be set for the script.
	PersistScript      bool                `json:"persistScript"`                // Specifies whether the compiled script should be persisted.
	ExecutionContextID *ExecutionContextID `json:"executionContextId,omitempty"` // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
}

// NewCompileScriptArgs initializes CompileScriptArgs with the required arguments.
func NewCompileScriptArgs(expression string, sourceURL string, persistScript bool) *CompileScriptArgs {
	args := new(CompileScriptArgs)
	args.Expression = expression
	args.SourceURL = sourceURL
	args.PersistScript = persistScript
	return args
}

// SetExecutionContextID sets the ExecutionContextID optional argument. Specifies in which execution context to perform script run. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
func (a *CompileScriptArgs) SetExecutionContextID(executionContextID ExecutionContextID) *CompileScriptArgs {
	a.ExecutionContextID = &executionContextID
	return a
}

// CompileScriptReply represents the return values for CompileScript in the Runtime domain.
type CompileScriptReply struct {
	ScriptID         *ScriptID         `json:"scriptId,omitempty"`         // Id of the script.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// EvaluateArgs represents the arguments for Evaluate in the Runtime domain.
type EvaluateArgs struct {
	Expression            string              `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           *string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	IncludeCommandLineAPI *bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	Silent                *bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	ContextID             *ExecutionContextID `json:"contextId,omitempty"`             // Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ReturnByValue         *bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	// GeneratePreview Whether preview should be generated for the result.
	//
	// Note: This property is experimental.
	GeneratePreview *bool `json:"generatePreview,omitempty"`
	UserGesture     *bool `json:"userGesture,omitempty"`  // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise    *bool `json:"awaitPromise,omitempty"` // Whether execution should `await` for resulting value and return once awaited promise is resolved.
}

// NewEvaluateArgs initializes EvaluateArgs with the required arguments.
func NewEvaluateArgs(expression string) *EvaluateArgs {
	args := new(EvaluateArgs)
	args.Expression = expression
	return args
}

// SetObjectGroup sets the ObjectGroup optional argument. Symbolic group name that can be used to release multiple objects.
func (a *EvaluateArgs) SetObjectGroup(objectGroup string) *EvaluateArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// SetIncludeCommandLineAPI sets the IncludeCommandLineAPI optional argument. Determines whether Command Line API should be available during the evaluation.
func (a *EvaluateArgs) SetIncludeCommandLineAPI(includeCommandLineAPI bool) *EvaluateArgs {
	a.IncludeCommandLineAPI = &includeCommandLineAPI
	return a
}

// SetSilent sets the Silent optional argument. In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (a *EvaluateArgs) SetSilent(silent bool) *EvaluateArgs {
	a.Silent = &silent
	return a
}

// SetContextID sets the ContextID optional argument. Specifies in which execution context to perform evaluation. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
func (a *EvaluateArgs) SetContextID(contextID ExecutionContextID) *EvaluateArgs {
	a.ContextID = &contextID
	return a
}

// SetReturnByValue sets the ReturnByValue optional argument. Whether the result is expected to be a JSON object that should be sent by value.
func (a *EvaluateArgs) SetReturnByValue(returnByValue bool) *EvaluateArgs {
	a.ReturnByValue = &returnByValue
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument. Whether preview should be generated for the result.
//
// Note: This property is experimental.
func (a *EvaluateArgs) SetGeneratePreview(generatePreview bool) *EvaluateArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// SetUserGesture sets the UserGesture optional argument. Whether execution should be treated as initiated by user in the UI.
func (a *EvaluateArgs) SetUserGesture(userGesture bool) *EvaluateArgs {
	a.UserGesture = &userGesture
	return a
}

// SetAwaitPromise sets the AwaitPromise optional argument. Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (a *EvaluateArgs) SetAwaitPromise(awaitPromise bool) *EvaluateArgs {
	a.AwaitPromise = &awaitPromise
	return a
}

// EvaluateReply represents the return values for Evaluate in the Runtime domain.
type EvaluateReply struct {
	Result           RemoteObject      `json:"result"`                     // Evaluation result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// GetPropertiesArgs represents the arguments for GetProperties in the Runtime domain.
type GetPropertiesArgs struct {
	ObjectID      RemoteObjectID `json:"objectId"`                // Identifier of the object to return properties for.
	OwnProperties *bool          `json:"ownProperties,omitempty"` // If true, returns properties belonging only to the element itself, not to its prototype chain.
	// AccessorPropertiesOnly If true, returns accessor properties (with getter/setter) only; internal properties are not
	// returned either.
	//
	// Note: This property is experimental.
	AccessorPropertiesOnly *bool `json:"accessorPropertiesOnly,omitempty"`
	// GeneratePreview Whether preview should be generated for the results.
	//
	// Note: This property is experimental.
	GeneratePreview *bool `json:"generatePreview,omitempty"`
}

// NewGetPropertiesArgs initializes GetPropertiesArgs with the required arguments.
func NewGetPropertiesArgs(objectID RemoteObjectID) *GetPropertiesArgs {
	args := new(GetPropertiesArgs)
	args.ObjectID = objectID
	return args
}

// SetOwnProperties sets the OwnProperties optional argument. If true, returns properties belonging only to the element itself, not to its prototype
// chain.
func (a *GetPropertiesArgs) SetOwnProperties(ownProperties bool) *GetPropertiesArgs {
	a.OwnProperties = &ownProperties
	return a
}

// SetAccessorPropertiesOnly sets the AccessorPropertiesOnly optional argument. If true, returns accessor properties (with getter/setter) only; internal properties are not
// returned either.
//
// Note: This property is experimental.
func (a *GetPropertiesArgs) SetAccessorPropertiesOnly(accessorPropertiesOnly bool) *GetPropertiesArgs {
	a.AccessorPropertiesOnly = &accessorPropertiesOnly
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument. Whether preview should be generated for the results.
//
// Note: This property is experimental.
func (a *GetPropertiesArgs) SetGeneratePreview(generatePreview bool) *GetPropertiesArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// GetPropertiesReply represents the return values for GetProperties in the Runtime domain.
type GetPropertiesReply struct {
	Result             []PropertyDescriptor         `json:"result"`                       // Object properties.
	InternalProperties []InternalPropertyDescriptor `json:"internalProperties,omitempty"` // Internal object properties (only of the element itself).
	ExceptionDetails   *ExceptionDetails            `json:"exceptionDetails,omitempty"`   // Exception details.
}

// GlobalLexicalScopeNamesArgs represents the arguments for GlobalLexicalScopeNames in the Runtime domain.
type GlobalLexicalScopeNamesArgs struct {
	ExecutionContextID *ExecutionContextID `json:"executionContextId,omitempty"` // Specifies in which execution context to lookup global scope variables.
}

// NewGlobalLexicalScopeNamesArgs initializes GlobalLexicalScopeNamesArgs with the required arguments.
func NewGlobalLexicalScopeNamesArgs() *GlobalLexicalScopeNamesArgs {
	args := new(GlobalLexicalScopeNamesArgs)

	return args
}

// SetExecutionContextID sets the ExecutionContextID optional argument. Specifies in which execution context to lookup global scope variables.
func (a *GlobalLexicalScopeNamesArgs) SetExecutionContextID(executionContextID ExecutionContextID) *GlobalLexicalScopeNamesArgs {
	a.ExecutionContextID = &executionContextID
	return a
}

// GlobalLexicalScopeNamesReply represents the return values for GlobalLexicalScopeNames in the Runtime domain.
type GlobalLexicalScopeNamesReply struct {
	Names []string `json:"names"` // No description.
}

// QueryObjectsArgs represents the arguments for QueryObjects in the Runtime domain.
type QueryObjectsArgs struct {
	PrototypeObjectID RemoteObjectID `json:"prototypeObjectId"` // Identifier of the prototype to return objects for.
}

// NewQueryObjectsArgs initializes QueryObjectsArgs with the required arguments.
func NewQueryObjectsArgs(prototypeObjectID RemoteObjectID) *QueryObjectsArgs {
	args := new(QueryObjectsArgs)
	args.PrototypeObjectID = prototypeObjectID
	return args
}

// QueryObjectsReply represents the return values for QueryObjects in the Runtime domain.
type QueryObjectsReply struct {
	Objects RemoteObject `json:"objects"` // Array with objects.
}

// ReleaseObjectArgs represents the arguments for ReleaseObject in the Runtime domain.
type ReleaseObjectArgs struct {
	ObjectID RemoteObjectID `json:"objectId"` // Identifier of the object to release.
}

// NewReleaseObjectArgs initializes ReleaseObjectArgs with the required arguments.
func NewReleaseObjectArgs(objectID RemoteObjectID) *ReleaseObjectArgs {
	args := new(ReleaseObjectArgs)
	args.ObjectID = objectID
	return args
}

// ReleaseObjectGroupArgs represents the arguments for ReleaseObjectGroup in the Runtime domain.
type ReleaseObjectGroupArgs struct {
	ObjectGroup string `json:"objectGroup"` // Symbolic object group name.
}

// NewReleaseObjectGroupArgs initializes ReleaseObjectGroupArgs with the required arguments.
func NewReleaseObjectGroupArgs(objectGroup string) *ReleaseObjectGroupArgs {
	args := new(ReleaseObjectGroupArgs)
	args.ObjectGroup = objectGroup
	return args
}

// RunScriptArgs represents the arguments for RunScript in the Runtime domain.
type RunScriptArgs struct {
	ScriptID              ScriptID            `json:"scriptId"`                        // Id of the script to run.
	ExecutionContextID    *ExecutionContextID `json:"executionContextId,omitempty"`    // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ObjectGroup           *string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	Silent                *bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides `setPauseOnException` state.
	IncludeCommandLineAPI *bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	ReturnByValue         *bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview       *bool               `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	AwaitPromise          *bool               `json:"awaitPromise,omitempty"`          // Whether execution should `await` for resulting value and return once awaited promise is resolved.
}

// NewRunScriptArgs initializes RunScriptArgs with the required arguments.
func NewRunScriptArgs(scriptID ScriptID) *RunScriptArgs {
	args := new(RunScriptArgs)
	args.ScriptID = scriptID
	return args
}

// SetExecutionContextID sets the ExecutionContextID optional argument. Specifies in which execution context to perform script run. If the parameter is omitted the
// evaluation will be performed in the context of the inspected page.
func (a *RunScriptArgs) SetExecutionContextID(executionContextID ExecutionContextID) *RunScriptArgs {
	a.ExecutionContextID = &executionContextID
	return a
}

// SetObjectGroup sets the ObjectGroup optional argument. Symbolic group name that can be used to release multiple objects.
func (a *RunScriptArgs) SetObjectGroup(objectGroup string) *RunScriptArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// SetSilent sets the Silent optional argument. In silent mode exceptions thrown during evaluation are not reported and do not pause
// execution. Overrides `setPauseOnException` state.
func (a *RunScriptArgs) SetSilent(silent bool) *RunScriptArgs {
	a.Silent = &silent
	return a
}

// SetIncludeCommandLineAPI sets the IncludeCommandLineAPI optional argument. Determines whether Command Line API should be available during the evaluation.
func (a *RunScriptArgs) SetIncludeCommandLineAPI(includeCommandLineAPI bool) *RunScriptArgs {
	a.IncludeCommandLineAPI = &includeCommandLineAPI
	return a
}

// SetReturnByValue sets the ReturnByValue optional argument. Whether the result is expected to be a JSON object which should be sent by value.
func (a *RunScriptArgs) SetReturnByValue(returnByValue bool) *RunScriptArgs {
	a.ReturnByValue = &returnByValue
	return a
}

// SetGeneratePreview sets the GeneratePreview optional argument. Whether preview should be generated for the result.
func (a *RunScriptArgs) SetGeneratePreview(generatePreview bool) *RunScriptArgs {
	a.GeneratePreview = &generatePreview
	return a
}

// SetAwaitPromise sets the AwaitPromise optional argument. Whether execution should `await` for resulting value and return once awaited promise is
// resolved.
func (a *RunScriptArgs) SetAwaitPromise(awaitPromise bool) *RunScriptArgs {
	a.AwaitPromise = &awaitPromise
	return a
}

// RunScriptReply represents the return values for RunScript in the Runtime domain.
type RunScriptReply struct {
	Result           RemoteObject      `json:"result"`                     // Run result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// SetCustomObjectFormatterEnabledArgs represents the arguments for SetCustomObjectFormatterEnabled in the Runtime domain.
type SetCustomObjectFormatterEnabledArgs struct {
	Enabled bool `json:"enabled"` // No description.
}

// NewSetCustomObjectFormatterEnabledArgs initializes SetCustomObjectFormatterEnabledArgs with the required arguments.
func NewSetCustomObjectFormatterEnabledArgs(enabled bool) *SetCustomObjectFormatterEnabledArgs {
	args := new(SetCustomObjectFormatterEnabledArgs)
	args.Enabled = enabled
	return args
}
