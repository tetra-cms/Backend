-- CreateTable
CREATE TABLE "Pages" (
    "id" SERIAL NOT NULL,
    "route" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "content" TEXT NOT NULL,

    CONSTRAINT "Pages_pkey" PRIMARY KEY ("id")
);
