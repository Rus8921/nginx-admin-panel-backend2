import { Dialog, DialogBackdrop, DialogPanel, DialogTitle, Listbox, ListboxButton, ListboxOption, ListboxOptions } from "@headlessui/react";
import { useContext, useState } from "react";
import WebsitesPageContext from "./WebsitesPageContext";
import { CommonButton } from "../../../ui/components/buttons/CommonButton";
import SectionBreak from "../../../ui/components/sections/SectionBreak";
import ContentBreak from "../../../ui/components/sections/ContentBreak";
import Input from "../../../ui/components/fields/Input";
import { ChevronDown, Terminal } from "react-feather";

function NewWebsiteModal() {
  let servers = [{id:1, name:"ExampleServer"}]
  const [selectedServer, setSelectedServer] = useState(servers[0])
  const context = useContext(WebsitesPageContext)

  return (
    <Dialog open={context.isNewWebsite} onClose={()=>{context.setIsNewWebsite(false)}} as="div" className="relative z-50 focus:outline-none">
      <DialogBackdrop transition className="fixed inset-0 w-screen bg-black/30" />

      <div className="fixed inset-0 w-screen overflow-y-auto flex min-h-full items-center justify-center p-12">
        <DialogPanel
          transition
          className="max-w-2xl rounded-xl bg-white p-12 flex flex-col gap-6 duration-300 ease-out data-[closed]:transform-[scale(95%)] data-[closed]:opacity-0 backdrop-blur-2xl"
        >
          <DialogTitle as="h2" >
            Add New Website
          </DialogTitle>
          <ContentBreak />
          <p>
            Fill the following form, and press 'Add' button to add new website for configuration.
          </p>
          <form className="flex flex-col gap-3 w-full *:w-full">
            <Input Icon={Terminal} placeholder="Enter website name" onChange={(e)=>{}}/>
            <Input Icon={Terminal} placeholder="Enter website domain" onChange={(e)=>{}}/>
            <div className="flex flex-row items-center w-[440px] relative">
              <ChevronDown className=" text-main-clr stroke-1 absolute ml-5" />
              <select name="server"
                className="bg-white pl-16 w-full border rounded-md border-scndry-txt-clr px-2 py-2.5 focus:border-main-clr focus-visible:border-main-clr outline-none appearance-none"
              >
                <option disabled>Select server to connect to</option>
                <option value="ex_server">Example Server</option>
              </select>
            </div>
              {/* <Listbox value={selectedServer} onChange={setSelectedServer}>
                <ListboxButton>{selectedServer.name}</ListboxButton>
                <ListboxOptions anchor="bottom">
                  {servers.map((server) => (
                    <ListboxOption key={server.id} value={server} className="data-[focus]:bg-blue-100">
                      {server.name}
                    </ListboxOption>
                  ))}
                </ListboxOptions>
              </Listbox> */}
            <div className="flex gap-4 mt-5">
              <CommonButton type="transparentBgMainText" onClick={()=>{context.setIsNewWebsite(false)}}
              >
                Cancel
              </CommonButton>
              <CommonButton isSubmit={true} type="blueBgWhiteText" onClick={()=>{context.setIsNewWebsite(false)}}
              >
                Add
              </CommonButton>
            </div>
          </form>
        </DialogPanel>
      </div>
  </Dialog>
  )
}

export default NewWebsiteModal;