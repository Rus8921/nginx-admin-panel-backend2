import { useContext, useEffect, useState } from "react";
import WebsitesPageContext from "./WebsitesPageContext";
import { CommonButton } from "../../../ui/components/buttons/CommonButton";
import Input from "../../../ui/components/fields/Input";
import { Globe, Hash } from "react-feather";
import { useNavigate } from "react-router-dom";
import { DefaultValues, useForm } from "react-hook-form";
import { IAllServersItem } from "../../../types/servers";
import nginxPanelApiService from "../../../api/NginxPanelApiService";

interface NewWebsiteFromFields {
  name: string,
  domain: string,
  // connectedServerId: number, 
}

function NewWebsiteForm() {
  const navigate = useNavigate();
  const context = useContext(WebsitesPageContext)

  const defaultValues:DefaultValues<NewWebsiteFromFields> = {
    name:"",
    domain:"",
    // connectedServerId:0
  }
  const { 
    handleSubmit,
    control,
    formState: {errors}
  } = useForm<NewWebsiteFromFields>({
    defaultValues:defaultValues, 
    mode: "onChange"
  });

  const onSubmit = handleSubmit((data) => {
    let id = context.websites.length
    context.setIsNewWebsite(false)
    navigate(`/websites/${id}`);
  });

  let [servers, setServers] = useState<IAllServersItem[]>([])
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    let data = nginxPanelApiService.getServers(0);
    data.then((resp) => {
      if (resp.status === 200) {
        setServers(resp.data.servers);
        setIsLoading(false);
      }
    });
  }, []);

  return (
          <form className="flex flex-col gap-3 w-full *:w-full" onSubmit={onSubmit}>
            <Input Icon={Hash} placeholder="Enter website name" type="text" control={control} name="name" rules={{required: true, validate: (value)=>
                  context.websites?.every((website)=> website.name!==value)
              }} />
            <Input 
              Icon={Globe} 
              placeholder="Enter website domain" 
              type="text" 
              control={control} 
              name="domain" 
              rules={{
                required: true, 
                validate: (value)=>
                  context.websites?.every((website)=> website.url!==value)
              }} 
            />
            <div className="text-scndry-clr">
              {(errors.name?.type==="required") && (<p>* fill the website name</p>)}
              {(errors.name?.type==="validate") && (<p>* website name is already taken</p>)}
              {(errors.domain?.type==="required") && (<p>* fill the website domain</p>)}
              {(errors.domain?.type==="validate") && (<p>* website domain is already taken</p>)}
              {/* {((errors.connectedServerId?.type==="required")||(errors.connectedServerId?.type==="min")||(errors.connectedServerId?.type==="disabled")) && (<p>Choose server to connect to</p>)} */}
            </div>
            <div className="flex gap-4 mt-5">
              <CommonButton type="transparentBgMainText" onClick={()=>{context.setIsNewWebsite(false)}}
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

export default NewWebsiteForm;