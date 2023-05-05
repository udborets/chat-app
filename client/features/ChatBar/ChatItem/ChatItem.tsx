import Image from 'next/image';
import { FC } from 'react';

import noAvatarWhite from './assets/no-avatar-white.png';
import { ChatItemProps } from './models';
import styles from './styles.module.scss';

const ChatItem: FC<ChatItemProps> = ({ companionAvatar, companionName, id, lastMessage }) => {
  return (
    <li className={`${styles.bgAnimation} flex w-[96%] h-[70px] min-h-[70px] text-[1.2rem] py-1 px-3 justify-center items-center gap-4 rounded-[10px]`}>
      <Image
        src={noAvatarWhite}
        alt='avatar'
        className="rounded-[50%] h-[45px] min-h-[45px] w-[45px] min-w-[45px] "
      />
      <div className="h-full flex w-full flex-col justify-around ">
        <span className="font-bold">
          {companionName}
        </span>
        <span>
          {lastMessage}
        </span>
      </div>
    </li>
  )
}

export default ChatItem;