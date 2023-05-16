import { FC, useState } from "react";

const SearchInput: FC = () => {
  const [searchQuery, setSearchQuery] = useState<string>('');
  return (
    <div className="w-full h-fit p-2">
      <input
        onChange={(e) => setSearchQuery(e.target.value)}
        value={searchQuery}
        type="text"
        placeholder="Enter search query..."
        className="outline outline-1 outline-gray-400 h-[30px] text-[1.1rem] w-full rounded-[5px] py-2 px-3"
      />
    </div>
  )
}

export default SearchInput;