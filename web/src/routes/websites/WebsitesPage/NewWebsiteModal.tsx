import { useContext } from "react";
import WebsitesPageContext from "./WebsitesPageContext";
import ContentBreak from "../../../ui/components/sections/ContentBreak";
import NewWebsiteForm from "./NewWebsiteForm";
import Modal from "../../../ui/components/modal/Modal";

function NewWebsiteModal() {
  const context = useContext(WebsitesPageContext)

  return (
    <Modal open={context.isNewWebsite} onClose={()=>{context.setIsNewWebsite(false)}} title="Add New Website">
      <ContentBreak />
      <p>
        Fill the following form, and press 'Add' button to add new website to configurate.
      </p>
      <NewWebsiteForm />
    </Modal>
  )
}

export default NewWebsiteModal;