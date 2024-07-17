import redis

r = redis.Redis(host="redis-service.learnk8s.svc.cluster.local", port=6379, db=0, decode_responses=True)

allUnpaidTaxes = r.hgetall("upaidtaxes")


def payTax(clientId: str, tax: float):
    print(f"Paid tax: {tax} for client: {clientId}")


for clientId, tax in allUnpaidTaxes.items():
    payTax(clientId, tax)
