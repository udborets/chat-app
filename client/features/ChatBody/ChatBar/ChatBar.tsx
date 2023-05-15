import { FC } from "react";
import Image from "next/image";

import manNoAvatar from '@/assets/images/manNoAvatar.png';

const ChatBar: FC = ({ }) => {
  return (
    <div
      className="px-4 flex gap-2 items-center"
    >
      <Image
        src={manNoAvatar}
        alt="Avatar image"
        width={50}
        height={50}
        className="w-[50px] h-[50px]"
      />
      <span
        className="font-bold"
      >
        Friend name
      </span>
      <span
        className="opacity-80"
      >
        was online 13.05.2023 15:23
      </span>
    </div>
  )
}

export default ChatBar;