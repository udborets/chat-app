import { FC, useRef } from "react"

const SearchBar: FC = () => {
  const searchRef = useRef<HTMLInputElement>(null);
  return (
    <input
      className="w-full h-[50px] bg-slate-300 p-2 text-[1.1rem] rounded-[10px]"
      placeholder="Search..."
      ref={searchRef}
    />
  )
}

export default SearchBar