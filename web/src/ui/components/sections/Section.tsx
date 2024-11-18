import { HTMLProps, PropsWithChildren } from "react";
import SectionBreak from "./SectionBreak";

function Section({ title, children, ...rest }: PropsWithChildren<HTMLProps<HTMLHeadingElement> & { title: string }>) {
  return (
    <section className="flex flex-col gap-6 w-full">
      <h3 {...rest}>{title}</h3>
      <SectionBreak />
      <div className="w-full flex flex-col gap-4  overflow-clip">
        {children}
      </div>
    </section>
  )
}

export default Section;