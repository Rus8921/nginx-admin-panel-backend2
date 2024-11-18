type TableItem<T> = T & {
  id: number,
}

export interface TableData {
  head?: {
    id: number,
    name: string
  },
  headings?: string[]
  items: TableItem<Object>[]
}

function FlatTable({ data }: { data: TableData[] }) {
  return (<table className="w-fit overflowx-x-scroll border-collapse text-scndry-txt-clr" >
    {data.map((itemSet) => (
      <>
        {!!itemSet.head && <thead key={itemSet.head.id}>
          <tr className="border-b border-scndry-txt-clr">
            <th className="text-left p-2 font-normal">{itemSet.head.name}</th>
          </tr>
        </thead>}
        {itemSet.items.map((item) => (
          <tr key={item.id} className="border-b border-scndry-txt-clr">
            {
              Object.entries(item).filter(entry => entry[0] !== "id").map(entry => (
                <td className="py-2 px-8 w-fit last:w-full whitespace-nowrap overflow-ellipsis">{entry[1]}</td>
              ))
            }
          </tr>
        ))}
      </>
    ))
    }
  </table >);
}

export default FlatTable;
