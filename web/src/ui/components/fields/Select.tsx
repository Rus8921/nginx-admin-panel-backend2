import { Icon } from "react-feather";
import { FieldPath, FieldValues, useController, UseControllerProps } from "react-hook-form";

function Select<T extends FieldValues>(props: React.ComponentProps<"select"> & UseControllerProps<T, FieldPath<T>> & { Icon: Icon;}) {
  const {name, rules, control, Icon, children, ...rest} = props
  const { field } = useController({
    name, 
    rules,
    control
  })

  return (
    <div className="flex flex-row items-center w-[440px] relative">
      <Icon className=" text-main-clr stroke-1 absolute ml-5" />
      <select defaultValue={0} className="bg-white pl-16 w-full border rounded-md border-scndry-txt-clr px-2 py-2.5 focus:border-main-clr focus-visible:border-main-clr outline-none appearance-none" {...field} {...rest}>
        {children}
      </select>
    </div>
  )
}

export default Select