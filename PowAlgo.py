import mpmath
import hashlib
#have multiple functions e.g. 30, take random combos of algos, check for the first x many digits = last digits of prev block passed through the algo.

typesOfFunc = ["powers", "logarithm", "trig", "hyperbolic", "exponentials", "factorials", "sha"]

##powers = random power of between -5, 5
##log base some no form -5, 5
trig = ["sin", "cos", "tan", "cot", "csc", "sec", "acos", "asin", "atan", "acsc"]
hyperbolic = ["cosh", "sinh", "tanh", "coth", "csch", "sech", "asinh", "atanh", "acosh", "acsch"]
##factorials n! - 1, 10 !
sha = ["sha265", "sha384", "sha224", "sha1", "md5"]
##final hash of sha 512
stdRange = (5, -5)


def powerFunc(number, power):
    return number**power

def logFunc(number, log):
    return log(number, log)

def trigFunc(number, funcNo):
    if number = 1:
        return(sin(number/maxNo)*maxNo)
    elif number = 2:
        return(cos(number/maxNo)*maxNo)
    elif number = 3:
        return(tan(number/maxNo)*maxNo)
    elif number = 4:
        return(cot(number/maxNo)*maxNo)
    elif number = 5:
        return(csc(number/maxNo)*maxNo)
    elif number = 6:
        return(sec(number/maxNo)*maxNo)
    elif number = 7:
        return(acos(number/maxNo)*maxNo)
    elif number = 8:
        return(asin(number/maxNo)*maxNo)
    elif number = 9:
        return(atan(number/maxNo)*maxNo)
    else number = 0:
        return(acsc(number/maxNo)*maxNo)


def factorialFunc(number, fact):
    return(factorial(number) - factorial(fact))

def shaFunc(number, shaNo):
    if shaNo = 1:
        return hashlib.sha256(number)
    elif shaNo = 2:
        return hashlib.sha384(number)

