import pandas as pd
import random
import faker

def generate_dataset(num_rows):
    # 创建一个空的 DataFrame，指定列名
    df = pd.DataFrame(columns=['url', 'intro', 'uname', 'password', 'email'])

    # 使用 Faker 库来生成虚假数据
    fake = faker.Faker()

    data_to_append = []  # 用于存储要追加的数据

    for _ in range(num_rows):
        # 生成虚假数据
        url = fake.url()
        intro = fake.sentence()
        uname = fake.user_name()
        password = fake.password()
        email = fake.email()

        data_to_append.append([url, intro, uname, password, email])

    # 使用 pd.concat() 将数据追加到 DataFrame
    df = pd.concat([df, pd.DataFrame(data_to_append, columns=df.columns)], ignore_index=True)

    return df

# 调用函数生成数据集，指定生成数据的条数

dataset = generate_dataset(1_000)
print(dataset.head())
dataset.to_csv(f"./dataset_0.1M.csv", index=False)
