import { Dialog, DialogBackdrop, DialogPanel, DialogTitle } from "@headlessui/react";
import { PropsWithChildren } from "react";

function Modal({open, onClose, title, children}:PropsWithChildren<{open:boolean, onClose:(value?:boolean)=>void, title:string}>) {
  return (
    <Dialog open={open} onClose={onClose} as="div" className="relative z-50 focus:outline-none">
      <DialogBackdrop transition className="fixed inset-0 w-screen bg-black/30" />

      <div className="fixed inset-0 w-screen overflow-y-auto flex min-h-full items-center justify-center p-12">
        <DialogPanel
          transition
          className="max-w-2xl rounded-xl bg-white p-12 flex flex-col gap-6 duration-300 ease-out data-[closed]:transform-[scale(95%)] data-[closed]:opacity-0 backdrop-blur-2xl"
        >
          <DialogTitle as="h2" >{title}</DialogTitle>
          {children}
        </DialogPanel>
      </div>
  </Dialog>
  )
}

export default Modal;