import { HTMLProps, PropsWithChildren } from "react";

function TextButton({ onClick, children }: PropsWithChildren<HTMLProps<HTMLButtonElement>>) {
  return (
    <button className="w-full text-right text-scndry-clr hover:text-red active:font-medium" onClick={onClick}>
      {children}
    </button>
  )
}

export default TextButton;