import redis

r = redis.Redis(
    host="redis-service.k8s-in-a-shell.svc.cluster.local",
    port=6379,
    db=0,
    decode_responses=True,
)

allUnpaidTaxes = r.hgetall("unpaidtaxes")
print(f"All unpaid taxes: {allUnpaidTaxes}")


def payTax(employeeId: str, tax: float):
    print(f"Paid tax: {tax} for employee: {employeeId}")


for employeeId, tax in allUnpaidTaxes.items():
    payTax(employeeId, tax)

r.delete("unpaidtaxes")
