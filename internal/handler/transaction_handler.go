package handler

import (
	"strconv"

	"github.com/aldisypu/go-pos/internal/domain"
	"github.com/aldisypu/go-pos/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var input domain.Transaction
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := h.transactionService.CreateTransaction(c.Context(), &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": input})
}

func (h *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := h.transactionService.GetAllTransactions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get transactions"})
	}

	return c.JSON(fiber.Map{"data": transactions})
}

func (h *TransactionHandler) GetTransactionByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	transaction, err := h.transactionService.GetTransactionByID(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Transaction not found"})
	}

	return c.JSON(fiber.Map{"data": transaction})
}

func (h *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	var updatedTransaction domain.Transaction
	if err := c.BodyParser(&updatedTransaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err = h.transactionService.UpdateTransaction(c.Context(), uint(id), &updatedTransaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update transaction"})
	}

	return c.JSON(fiber.Map{"data": updatedTransaction})
}

func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	err = h.transactionService.DeleteTransaction(c.Context(), uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete transaction"})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
