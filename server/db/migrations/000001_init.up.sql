CREATE TABLE "auth" (
                               "user_id" integer NOT NULL UNIQUE,
                               "name" varchar(40) NOT NULL UNIQUE,
                               "email" varchar(30) UNIQUE,
                               "phone" varchar(20) UNIQUE,
                               "hash_password" varchar(255) NOT NULL
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "users" (
                                "id" serial NOT NULL,
                                "name" varchar(20) NOT NULL UNIQUE,
                                "hash_password" varchar(255) NOT NULL,
                                "email" varchar(30) UNIQUE,
                                "phone" varchar(20) UNIQUE,
                                "avatar_url" varchar(255),
                                "last_seen" integer,
                                "created_at" integer NOT NULL,
                                "updated_at" integer,
                                CONSTRAINT "users_pk" PRIMARY KEY ("id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "chats" (
                                "chat_id" serial NOT NULL,
                                "last_message_id" integer,
                                "created_at" integer,
                                "updated_at" integer,
                                CONSTRAINT "chats_pk" PRIMARY KEY ("chat_id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "messages" (
                                   "message_id" serial NOT NULL,
                                   "chat_id" integer NOT NULL,
                                   "text" varchar(255) NOT NULL,
                                   "sender_id" integer NOT NULL,
                                   "receiver_id" integer NOT NULL,
                                   "is_seen" BOOLEAN NOT NULL,
                                   "created_at" integer,
                                   "updated_at" integer,
                                   CONSTRAINT "messages_pk" PRIMARY KEY ("message_id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "users_chats" (
                                      "user_id" integer NOT NULL,
                                      "chat_id" integer NOT NULL,
                                      CONSTRAINT "users_chats_pk" PRIMARY KEY ("user_id","chat_id")
) WITH (
      OIDS=FALSE
      );



CREATE TABLE "friends" (
                                  "user_id" integer NOT NULL,
                                  "friend_id" integer NOT NULL,
                                  CONSTRAINT "friends_pk" PRIMARY KEY ("user_id")
) WITH (
      OIDS=FALSE
      );



ALTER TABLE "auth" ADD CONSTRAINT "auth_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id");


ALTER TABLE "chats" ADD CONSTRAINT "chats_fk0" FOREIGN KEY ("last_message_id") REFERENCES "messages"("message_id");

ALTER TABLE "messages" ADD CONSTRAINT "messages_fk0" FOREIGN KEY ("chat_id") REFERENCES "chats"("chat_id");
ALTER TABLE "messages" ADD CONSTRAINT "messages_fk1" FOREIGN KEY ("sender_id") REFERENCES "users"("id");
ALTER TABLE "messages" ADD CONSTRAINT "messages_fk2" FOREIGN KEY ("receiver_id") REFERENCES "users"("id");

ALTER TABLE "users_chats" ADD CONSTRAINT "users_chats_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id");
ALTER TABLE "users_chats" ADD CONSTRAINT "users_chats_fk1" FOREIGN KEY ("chat_id") REFERENCES "chats"("chat_id");

ALTER TABLE "friends" ADD CONSTRAINT "friends_fk0" FOREIGN KEY ("user_id") REFERENCES "users"("id");






