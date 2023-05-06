import { FC, useRef } from "react"

import styles from './styles.module.scss';

const SearchBar: FC = () => {
  const searchRef = useRef<HTMLInputElement>(null);
  return (
    <input
      className={`${styles.searchBar} w-full h-[50px] min-h-[50px] message-bg px-5 text-[1.15rem] rounded-[30px]`}
      placeholder="Search..."
      ref={searchRef}
    />
  )
}

export default SearchBar;