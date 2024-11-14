import { ComponentProps } from "react";

type TextButtonPropsWrapper<T> = T & {
  isDanger: boolean
}

function TextButton({ isDanger, children, className, ...rest }: TextButtonPropsWrapper<ComponentProps<"button">>) {
  return (
    <button className={"w-fit text-scndry-clr active:font-medium hover:underline " + (isDanger ? "hover:text-red " : "hover:text-main-clr ") + className} {...rest}>
      {children}
    </button >
  )
}

export default TextButton;