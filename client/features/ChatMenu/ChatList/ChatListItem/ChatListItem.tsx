import { FC } from "react";
import Image from "next/image";

import { ChatListItemProps } from "./models";
import manNoAvatar from '@/assets/images/manNoAvatar.png';

const ChatListItem: FC<ChatListItemProps> = ({ }) => {
  return (
    <li className="flex justify-between w-full px-2 py-1 hover:bg-gray-200 transition-colors duration-150">
      <div className="flex gap-2">
        <Image
          src={manNoAvatar}
          alt="Avatar icon"
          className="w-[45px] h-[45px]"
          width={45}
          height={45}
        />
        <div className="flex flex-col justify-between">
          <span className="font-bold">
            Name
          </span>
          <span>
            Last message
          </span>
        </div>
      </div>
      <span className="opacity-80">
        15:32
      </span>
    </li>
  )
}

export default ChatListItem