// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {repository} from '../models';
import {model} from '../models';

export function Create(arg1:repository.CreateDataModel):Promise<void>;

export function Delete(arg1:string):Promise<void>;

export function DeleteVariable(arg1:string,arg2:string):Promise<void>;

export function GetModelInfo(arg1:string):Promise<Array<model.ModelInfo>>;

export function GetVariableList(arg1:string):Promise<Array<model.Variable>>;

export function List(arg1:string,arg2:number,arg3:number):Promise<Array<model.DataModel>>;

export function Modify(arg1:repository.RequestModifyDataModel):Promise<void>;

export function RunCode(arg1:string):Promise<void>;

export function SaveModelInfo(arg1:string,arg2:Array<repository.RequestModelInfoSave>):Promise<void>;

export function SaveVariable(arg1:string,arg2:Array<model.VariableBase>):Promise<void>;

export function TestRunCode(arg1:string,arg2:string):Promise<void>;
