def div_and_remainder(n,d):  # 関数の定義の始まり
    if d == 0:
        raise Exception("0での除算はできません")
    return n / d, n % d      # 関数の定義終わり

v = div_and_remainder(5,2)   # mainの始まり
print(v)  # (2.5, 1)
