import openai
import os
from dotenv import load_dotenv

load_dotenv()
openai.api_key=os.getenv("OPENAI_API_KEY")

if __name__ == '__main__':
    message = [
        {"role": "system", "content": "You are a helpful assistant."},
    ]

    while True:
        prompt = input("Please input your query: ")
        message.append({
            "role":"user", "content":prompt,
        })
        response = openai.ChatCompletion.create(model="gpt-3.5-turbo",messages=message,temperature=0)

        content = response['choices'][0]['message']['content']
        print(content)
        message.append({
            "role":"assistant","content":content,
        })
