import { FC } from "react";

import Chat from "@/features/Chat/Chat";
import type { ChatProps } from "@/features/Chat/models";
import ChatBar from "@/features/ChatBar/ChatBar";
import type { ChatBarProps } from "@/features/ChatBar/models";
import PageLayout from "@/layouts/PageLayout/PageLayout";

const testChats: ChatBarProps = {
  chats: [
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random().toString(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 2",
      id: Math.random().toString(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 3",
      id: Math.random().toString(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 4",
      id: Math.random().toString(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 5",
      id: Math.random().toString(),
      lastMessage: "Hello!",
    },
  ]
}

const chatProps: ChatProps = {
  messages: [
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "",
      sendingTime: Date.now(),
      isRead: false,
    }, {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "",
      sendingTime: Date.now(),
      isRead: false,
    }, {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: false,
      text: "skjfs",
      sendingTime: Date.now(),
      isRead: true,
    },
    {
      id: Math.random().toString(),
      isOwn: true,
      text: "",
      sendingTime: Date.now(),
      isRead: false,
    },
  ]
}

const ChatPage: FC = () => {
  return (
    <PageLayout title="Chat" className="w-full h-full flex">
      <ChatBar {...testChats} />
      <Chat messages={chatProps.messages} />
    </PageLayout>
  )
}

export default ChatPage;