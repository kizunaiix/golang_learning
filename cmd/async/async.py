from asyncio import sleep, run
import time

async def my_async_func():
    time.sleep(3)
    print("wake!")


print("start")
run(my_async_func())
print("end.")
