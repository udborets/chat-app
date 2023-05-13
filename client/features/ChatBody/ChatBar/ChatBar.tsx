import { FC } from "react";
import Image from "next/image";

import manNoAvatar from '@/assets/manNoAvatar.png';

const ChatBar: FC = ({ }) => {
  return (
    <div
      className="flex gap-2"
    >
      <Image
        src={manNoAvatar}
        alt="Avatar image"
        width={70}
        height={70}
        className="w-[70px] h-[70px]"
      />
      <span>
        Companion name
      </span>
    </div>
  )
}

export default ChatBar;