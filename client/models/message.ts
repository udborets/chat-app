export type Message = {
  message_id: number;
  chat_id: number;
  text: string;
  sender_id: number;
  receiver_id: number;
  is_seen: boolean;
  created_at: number;
  updated_at: number;
};
