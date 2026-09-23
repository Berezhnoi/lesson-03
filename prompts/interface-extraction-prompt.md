# Промпт для виокремлення інтерфейсу (Розділ 2)

Вставте нижче точний текст промпту, який ви розробили для Розділу 2
домашньої роботи. Промпт має явно містити:

- **роль** ШІ (наприклад, "Ти — досвідчений Go-розробник, що
  спеціалізується на чистій архітектурі...");
- **контекст** — наведений у завданні код `OrderService`;
- **завдання** — виокремити мінімальний інтерфейс `OrderStore`,
  переписати `OrderService` на dependency injection, згенерувати мок;
- **обмеження** — маленький інтерфейс (приказка Роба Пайка),
  конкретний стиль коду, тощо;
- **бажаний формат виводу** (наприклад, "поверни лише Go-код у трьох
  блоках: інтерфейс, сервіс, мок").

<!-- TODO(розділ-2): видаліть цей HTML-коментар і вставте текст
     вашого промпту нижче цього рядка, звичайним видимим текстом
     Markdown. -->

## Мій промпт

### Role

You are an experienced Go developer specializing in clean architecture, dependency injection, and testable code.

### Context

The repository already contains `OrderService` in `orders/orders.go`. It is partially prepared for refactoring. Inspect the existing code and tests before making changes.

### Task

Refactor `OrderService` so that:

1. It depends on a minimal `OrderStore` interface instead of `*sql.DB`.
2. `OrderService` receives the store through `NewOrderService(store OrderStore)`.
3. `PlaceOrder` uses the store to execute the order query and correctly returns any error.
4. Add a simple mock/fake implementation of `OrderStore` for testing without a real database.
5. Create a separate `orders_ai_test.go` file containing tests for the refactored `OrderService`. The tests should cover at least:

   * a successful `PlaceOrder` call;
   * an error returned by the store;
   * that the store's `Exec` method is called with the expected arguments.

### Constraints

* Keep `OrderStore` as small as possible. Follow Rob Pike's principle: **"The bigger the interface, the weaker the abstraction."**
* Keep the implementation simple and idiomatic Go.
* Do not add unnecessary abstractions or dependencies.
* Do not change unrelated code.
* Do not use a real database in `orders_ai_test.go`.
* Keep the AI-generated tests separate from the existing tests.
* Run `gofmt` on all modified Go files.
* Make sure all existing and new tests pass.

### Expected output

Make the changes directly in the repository.

Create or modify the necessary files, including:

* `orders/orders.go` — refactored interface and service.
* `orders/orders_ai_test.go` — mock/fake and tests generated for the refactored service.

After making the changes, briefly report what was changed and the result of running the tests.
