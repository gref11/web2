-- 1
select emp.last_name, dept.department_name from
employee emp, department dept
where emp.deptartment_id = dept.deptartment_id;

-- 2
select * from
employee emp, department dept
where emp.deptartment_id = 60 and dept.deptartment_id = 60;

-- 3
select emp.employee_id, emp.manager from
employee emp;

-- 4
select dept.department_name, count(emp.employee_id) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name;

-- 5
select dept.department_name, count(emp.employee_id) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name
having count(emp.employee_id) > 5;

-- 6
select dept.department_name, avg(emp.salary) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name;


-- 7
select emp.employee_id
from employee emp, (select select dept.department_id d_id, avg(emp.salary) d_sal
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_id;) dept_avg_sal
where emp.department_id = dept_avg_sal.d_id and emp.salary > dept_avg_sal.d_sal;

-- 8
select *
from employee emp
where emp.manager = null;

-- 9
select 

-- 10
select unique emp.job
from employee emp;


insert into lang(lang_name) values ('');
insert into app_ability (app_id, lang_id) values ();
insert into app_ability (app_id, lang_id) select ?, lang_id from lang where lang_name = ? limit 1;

select name, lang_name
from application app 
join app_ability app_ab on app.app_id = app_ab.app_id
join lang on app_ab.lang_id = lang.lang_id;




-- 1
select emp.employee_id
from HR.employees emp
join (select dept.department_id, avg(emp1.salary) d_sal
from HR.employees emp1 
join HR.departments dept on emp1.department_id = dept.department_id
group by dept.department_id) dept_avg_sal
on emp.department_id = dept_avg_sal.department_id
where emp.salary > dept_avg_sal.d_sal;

-- 2
select *
from HR.employees
where department_id = (
    select department_id
    from hr.departments
    group by department_id
    order by count(*) desc
    fetch first 1 row only
);

-- 3
select *
from hr.employees
where salary > (
    select salary
    from hr.employees
    where employee_id = 101
    fetch first 1 row only
);

-- 4
select * 
from hr.employees
where department_id = (
    select department_id
    from hr.employees
    where employee_id = 102
    fetch first 1 row only
);

-- 5
select *
from hr.employees
where salary > (
    select max(salary)
    from hr.employees
    where department_id = 30
);

-- 6
select *
from hr.employees emp
join (
    select job_id, avg(salary) avg_salary
    from hr.employees
    group by job_id
) avg_sal on emp.job_id = avg_sal.job_id
where e.salary < avg_sal.avg_salary;

-- 7
select *
from


insert into user_auth(user_id, login, password_hash) values (?, ?, ?);