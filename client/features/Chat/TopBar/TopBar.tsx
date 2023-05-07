import { observer } from "mobx-react-lite";
import Image from "next/image";

import { chatBarState } from "@/store/ChatBarState";
import menuIconWhite from './assets/menuIconWhite.png';

const TopBar = observer(() => {
  return (
    <div className="w-full h-[60px] min-h-[40px] flex bg-secondary-lighter px-4 items-center">
      <button
        onClick={() => chatBarState.toggleIsActive()}
        className="pc:hidden h-fit w-fit p-3">
        <Image
          className="w-[20px] h-[20px]"
          src={menuIconWhite}
          alt='menu icon'
        />
      </button>
    </div>
  )
})

export default TopBar;