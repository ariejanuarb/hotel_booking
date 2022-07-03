menambah kolom "invoice id" di table discount
ALTER TABLE `discount` ADD `invoice_id` INT NOT NULL AFTER `discount_amount`;
ganti nama di table discount dari "discount_ammount" menjadi "discount_amount"

ganti event start dan event end dari time menjadi string biar bisa di insert sesuatu didalamnya