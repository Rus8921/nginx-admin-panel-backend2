import { useContext } from "react";
import ContentBreak from "../../../ui/components/sections/ContentBreak";
// import NewWebsiteForm from "./NewWebsiteForm";
import Modal from "../../../ui/components/modal/Modal";
import ServersPageContext from "./ServersPageContext";
import NewServerForm from "./NewServerForm";

function NewServerModal() {
  const context = useContext(ServersPageContext)

  return (
    <Modal open={context.isNewServer} onClose={()=>{context.setIsNewServer(false)}} title="Add New Server">
      <ContentBreak />
      <p>
        Fill the following form, and press 'Add' button to add new server to configurate.
      </p>
      <NewServerForm />
    </Modal>
  )
}

export default NewServerModal;