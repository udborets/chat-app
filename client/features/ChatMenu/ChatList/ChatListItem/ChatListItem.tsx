import { FC } from "react";
import Image from "next/image";

import { ChatListItemProps } from "./models";
import manNoAvatar from '@/assets/images/manNoAvatar.png';

const ChatListItem: FC<ChatListItemProps> = ({ }) => {
  return (
    <li>
      <div>
        <Image
          src={manNoAvatar}
          alt="Avatar icon"
        />
      </div>
    </li>
  )
}

export default ChatListItem