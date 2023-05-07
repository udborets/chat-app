import Image from "next/image";
import { FC, useRef, useState } from "react";

import crossIconWhite from './assets/crossIconWhite.png';
import searchIconWhite from './assets/searchIconWhite.png';

const SearchBar: FC = () => {
  const searchRef = useRef<HTMLInputElement>(null);
  const [searchQuery, setSearchQuery] = useState<string>('');
  const [isFocused, setIsFocused] = useState<boolean>(false);
  const focus = () => {
    searchRef.current?.focus();
    setIsFocused(true);
  }
  const blur = () => {
    searchRef.current?.blur();
    setIsFocused(false);
  }
  const clearSearchQuery = () => {
    setSearchQuery('');
  }
  return (
    <div
      onClick={focus}
      className={`duration-200 transition-all ease-out ${isFocused
        ? "outline-[var(--color-bg)] hover:outline-[var(--color-bg)]"
        : "outline-[rgba(255,255,255,0.34)] hover:outline-[rgba(255,255,255,0.62)]"
        } w-full h-[40px] min-h-[40px] message-bg px-[15px] rounded-[30px] outline-2 outline flex items-center gap-4`}
    >
      <Image
        className="h-[15px] w-[15px]"
        src={searchIconWhite}
        alt="Search icon"
      />
      <input
        className="h-full w-full outline-none message-bg text-[1rem]"
        onFocus={focus}
        onBlur={blur}
        placeholder="Search..."
        ref={searchRef}
        value={searchQuery}
        onChange={(e) => setSearchQuery(e.target.value)}
      />
      {searchQuery !== ''
        ? <button
          onClick={clearSearchQuery}
          className={` w-fit h-fit p-1 hover:bg-slate-800 transition-all duration-200 rounded-[50%]`}
        >
          <Image
            className={`h-[25px] w-[25px] min-h-[25px] min-w-[25px]`}
            src={crossIconWhite}
            alt="Cross icon"
          />
        </button>
        : ''}
    </div>
  )
}

export default SearchBar;