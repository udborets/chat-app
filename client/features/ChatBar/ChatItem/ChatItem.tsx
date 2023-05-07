import { observer } from 'mobx-react-lite';
import Image from 'next/image';
import { FC } from 'react';

import { currentChat } from '@/store/CurrentChat';
import noAvatarWhite from './assets/no-avatar-white.png';
import { ChatItemProps } from './models';
import styles from './styles.module.scss';

const ChatItem: FC<ChatItemProps> = ({ companionAvatar, companionName, id, lastMessage }) => {
  return (
    <li
      onClick={() => currentChat.setCurrentChat({ ...currentChat.getSelf(), id: id })}
      className={`${styles.chatItem} ${currentChat.id === id ? "bg-color" : "hover:bg-secondary-hover"}`}>
      <Image
        src={noAvatarWhite}
        alt='avatar'
        className="rounded-[50%] h-[40px] min-h-[40px] w-[40px] min-w-[40px] "
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

export default observer(ChatItem);