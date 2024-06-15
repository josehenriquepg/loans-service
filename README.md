# Loans service

Project for a service that determines what types of loans a person has access to.
The challenge proposes that we built a web API that will receive a POST request with a customer data we should follow some rules.

**Challenge:** [Loans Challenge](https://github.com/backend-br/desafios/blob/master/loans/PROBLEM.md) 
**Offered by:** [Back-End Brasil](https://github.com/backend-br)

## Rules

* Grant the personal loan if the client's salary is equal to or less than R$ 3000.
* Grant the personal loan if the client's salary is between R$3000 and R$5000, and the client is under 30 years old and resides in São Paulo(SP).
* Grant the payroll loan if the client's salary is equal to or greater than R$5000.
* Grant the loan with collateral if the client's salary is equal to or less than R$3000.
* Grant the loan with collateral if the client's salary is between R$3000 and R$5000, and the client is under 30 years old and resides in São Paulo(SP).

## Example Payload Response

```json
{
  "customer": "Vuxaywua Zukiagou",
  "loans": [
    {
      "type": "PERSONAL",
      "interest_rate": 4
    },
    {
      "type": "GUARANTEED",
      "interest_rate": 3
    },
    {
      "type": "CONSIGNMENT",
      "interest_rate": 2
    }
  ]
}
```
## 💻 Technologies used
[![techs](https://skillicons.dev/icons?i=go,git&theme=dark)](https://skillicons.dev)
## 🤝 Collaborators 
<table>
  <tr>
    <td align="center">
      <a href="http://github.com/josehenriquepg">
        <img src="https://avatars.githubusercontent.com/josehenriquepg" width="100px;" alt="Foto de José Henrique no GitHub"/><br>
        <sub>
          <b>José Henrique</b>
        </sub>
      </a>
    </td>
  </tr>
</table>