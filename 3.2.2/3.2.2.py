def maior_25(x):
    if (x > 25):
        return True


def ffilter(f, l):
    resultado = []
    if len(l) == 0:
        return resultado
    for i in range(len(l)):
        if (f(l[i]) is True):
            resultado.append(l[i])
            ffilter(f, l[(i+1):])
    return resultado


