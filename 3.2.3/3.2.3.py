def triplica(x):
    return x*3


def fmap(f, l):
    if len(l) == 0:
        return l
    for i in range(len(l)):
        l[i] = f(l[i])
    return l

   
