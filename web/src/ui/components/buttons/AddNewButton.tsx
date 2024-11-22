import { ComponentProps } from "react";
import { AddItemTargets } from "../../../types";
import { Plus } from "react-feather";

function AddNewButton({ target, className, ...rest }: ComponentProps<"button"> & { target: AddItemTargets }) {
  return (
    <button className="w-full h-full p-8 bg-transparent rounded-2xl flex flex-col items-center text-center justify-center gap-5 text-scndry-txt-clr hover:text-main-clr active:text-scndry-clr dashed-border" {...rest}>
      <Plus size={80} strokeWidth={1}/>
      <h3 className="text-inherit">Add New {target}</h3>
    </button>
  )
}

export default AddNewButton;