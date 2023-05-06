import Image from "next/image";
import { FC, useRef, useState } from "react";

import crossIconWhite from './assets/crossIconWhite.png';
import searchIconWhite from './assets/searchIconWhite.png';
import styles from './styles.module.scss';

const SearchBar: FC = () => {
  const searchRef = useRef<HTMLInputElement>(null);
  const [isFocused, setIsFocused] = useState<boolean>(false);
  const focus = () => {
    searchRef.current?.focus();
    setIsFocused(true);
  }
  const blur = () => {
    searchRef.current?.blur();
    setIsFocused(false);
  }
  return (
    <div
      onClick={(e) => {
        focus()
      }}
      className={`${styles.searchBar} ${isFocused
        ? "outline-[var(--color-bg)] hover:outline-[var(--color-bg)]"
        : "outline-[rgba(255,255,255,0.34)] hover:outline-[rgba(255,255,255,0.62)]"
        } w-full h-[50px] min-h-[50px] message-bg px-[15px] rounded-[30px] outline-2 outline flex items-center gap-4`}
    >
      <Image
        className="h-[25px] w-[25px]"
        src={searchIconWhite}
        alt="Search icon"
      />
      <input
        className="h-full w-full outline-none message-bg text-[1.15rem]"
        onFocus={focus}
        onBlur={blur}
        placeholder="Search..."
        ref={searchRef}
      />
      <Image
        onClick={() => {
          if (searchRef.current && searchRef.current.value)
            searchRef.current.value = '';
        }}
        className={`${isFocused
          ? "opacity-100 rotate-0"
          : "opacity-0 rotate-90"} hover:bg-slate-800 p-[2px] h-[35px] w-[35px] rounded-[50%] transition-all duration-200`}
        src={crossIconWhite}
        alt="Cross icon"
      />
    </div>
  )
}

export default SearchBar;