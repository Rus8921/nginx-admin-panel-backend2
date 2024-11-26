import { useContext } from "react";
import { CommonButton } from "../../../ui/components/buttons/CommonButton";
import Input from "../../../ui/components/fields/Input";
import { Globe, Hash, Terminal } from "react-feather";
import { useNavigate } from "react-router-dom";
import { DefaultValues, useForm } from "react-hook-form";
import ServersPageContext from "./ServersPageContext";

interface NewServerFormFields {
  name: string,
  domain: string, 
  ip: string
}

function NewServerForm() {
  const navigate = useNavigate();
  const context = useContext(ServersPageContext)

  const defaultValues:DefaultValues<NewServerFormFields> = {
    name:"",
    domain:"",
    ip: ""
  }
  const { 
    handleSubmit,
    control,
    formState: {errors}
  } = useForm<NewServerFormFields>({
    defaultValues:defaultValues, 
    mode: "onChange"
  });

  const onSubmit = handleSubmit((data) => {
    let id = context.servers.length
    context.setIsNewServer(false)
    navigate(`/servers/${id}`);
  });

  return (
          <form className="flex flex-col gap-3 w-full *:w-full" onSubmit={onSubmit}>
            <Input Icon={Hash} placeholder="Enter server name" type="text" control={control} name="name" rules={{required: true, validate: (value)=>
                  context.servers?.every((server)=> server.serverName!==value)
              }} />
            <Input 
              Icon={Globe} 
              placeholder="Enter server domain" 
              type="text" 
              control={control} 
              name="domain" 
              rules={{
                required: true, 
                validate: (value)=>
                  context.servers?.every((server)=> server.domainName!==value)
              }} 
            />
            <Input 
              Icon={Terminal} 
              placeholder="Enter server IP address" 
              type="text" 
              control={control} 
              name="ip" 
              rules={{
                required: true, 
                pattern: /^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$/,
                validate: (value)=>
                  context.servers?.every((server)=> server.ip!==value)
              }} 
            />
            <div className="text-scndry-clr">
              {(errors.name?.type==="required") && (<p role="alert">* fill the server name</p>)}
              {(errors.name?.type==="validate") && (<p role="alert">* this name is already taken</p>)}
              {(errors.domain?.type==="required") && (<p role="alert">* fill the server domain</p>)}
              {(errors.domain?.type==="validate") && (<p role="alert">* this domain is already taken</p>)}
              {(errors.ip?.type==="required") && (<p role="alert">* fill the server IP address</p>)}
              {(errors.ip?.type==="validate") && (<p role="alert">* this IP address is already taken</p>)}
              {(errors.ip?.type==="pattern") && (<p role="alert">* invalid IP address</p>)}
            </div>
            <div className="flex gap-4 mt-5">
              <CommonButton type="transparentBgMainText" onClick={()=>{context.setIsNewServer(false)}}
              >
                Cancel
              </CommonButton>
              <CommonButton isSubmit={true} type="blueBgWhiteText" onClick={()=>{}}
              >
                Add
              </CommonButton>
            </div>
          </form>
  )
}

export default NewServerForm;