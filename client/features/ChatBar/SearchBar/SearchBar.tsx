import { FC, useRef } from "react"

import styles from './styles.module.scss';

const SearchBar: FC = () => {
  const searchRef = useRef<HTMLInputElement>(null);
  return (
    <input
      className={`${styles.inputOutline} w-full h-[40px] min-h-[40px] message-bg px-5 rounded-[30px]`}
      placeholder="Search..."
      ref={searchRef}
    />
  )
}

export default SearchBar;