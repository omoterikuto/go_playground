SELECT Hell.id, Hell.name, el.name AS element, Grade.name AS grade
from ElementCompatibility AS e
    left join Element AS el ON e.element_id = el.id
    left join Element AS wel ON e.weakness_element_id = wel.id
    left join Hell ON Hell.element_id = e.id
    left join Grade ON Hell.grade_id = grade.id
WHERE wel.name = 'Graffiacane' && Grade.name = 'Boss'