import asyncio
import time

async def my_async_func(seconds: int):
    await asyncio.sleep(seconds)
    print(f"wake after {seconds} !")

async def main():
    print(f"[{time.asctime()}] start")

    # 并发运行两个 my_async_func 函数
    await asyncio.gather(
        my_async_func(3),
        my_async_func(5),
    )

    print(f"[{time.asctime()}] end.")

# 运行异步主函数
asyncio.run(main())
